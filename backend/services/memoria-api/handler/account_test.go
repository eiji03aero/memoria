package handler

import (
	"net/http"
	"testing"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/service"
	"memoria-api/domain/value"
	"memoria-api/testutil"
	"memoria-api/usecase"
	"memoria-api/util"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountSignup_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		name          string
		email         string
		userSpaceName string
		password      string
	}{
		{
			name:          "Hoge Tarou",
			email:         "hoge@gmail.com",
			userSpaceName: "Hoge family",
			password:      "password",
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"name":            util.StrToNilIfEmpty(test.name),
				"email":           util.StrToNilIfEmpty(test.email),
				"user_space_name": util.StrToNilIfEmpty(test.userSpaceName),
				"password":        util.StrToNilIfEmpty(test.password),
			},
		)

		// setup assertion
		api.MockMailer.
			EXPECT().
			Send(
				gomock.Any(),
				gomock.Cond(func(dto any) bool {
					return (dto).(interfaces.MailerSendDTO).To[0] == test.email
				}),
			)

		// -------------------- execution --------------------
		accountH := NewAccount()
		status, data, err := accountH.Signup(c, reg)

		// -------------------- assertion --------------------
		authUc := usecase.NewAuth(reg)
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := AccountSignupRes{}
		verifyJSONEncoding(data, &decodedRes)

		ret, err := authUc.VerifyJWT(decodedRes.Token)
		assert.NoError(t, err)
		assert.NotEmpty(t, ret.UserID)
		assert.NotEmpty(t, ret.UserSpaceID)

		// user record created
		users, err := reg.NewUserRepository().Find(&repository.FindOption{})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))
		user := users[0]
		assert.Equal(t, test.name, user.Name)
		assert.Equal(t, test.email, user.Email)
		assert.Equal(t, value.UserAccountStatus_Invited, user.AccountStatus)
		matched, err := service.CompareHashAndPassword([]byte(user.PasswordHash), []byte(test.password))
		assert.NoError(t, err)
		assert.True(t, matched)
		// user space created
		userSpaces, err := reg.NewUserSpaceRepository().Find(&repository.FindOption{})
		assert.Equal(t, 1, len(userSpaces))
		userSpace := userSpaces[0]
		assert.Equal(t, test.userSpaceName, userSpace.Name)
		// uusr record created
		uusrs, err := reg.NewUserUserSpaceRelationRepository().Find(&repository.FindOption{})
		assert.Equal(t, 1, len(uusrs))
		uusr := uusrs[0]
		assert.Equal(t, user.ID, uusr.UserID)
		assert.Equal(t, userSpace.ID, uusr.UserSpaceID)
		// user invitation record created
		uis, err := reg.NewUserInvitationRepository().Find(&repository.FindOption{})
		assert.Equal(t, 1, len(uis))
		ui := uis[0]
		assert.Equal(t, user.ID, ui.UserID)
		assert.Equal(t, value.UserInvitationType_Signup, ui.Type)
	}
}

func TestAccountSignup_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		name           string
		email          string
		userSpaceName  string
		password       string
		validationKey  cerrors.ValidationKey
		validationName string
		beforeRun      func()
	}{
		{
			name:           "",
			email:          "hoge@gmail.com",
			userSpaceName:  "Hoge family",
			password:       "password",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "name",
		},
		{
			name:           "Hoge",
			email:          "",
			userSpaceName:  "Hoge family",
			password:       "password",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "email",
		},
		{
			name:           "Hoge",
			email:          "hmm",
			userSpaceName:  "Hoge family",
			password:       "password",
			validationKey:  cerrors.ValidationKey_InvalidFormat,
			validationName: "email",
		},
		{
			name:           "Hoge",
			email:          "hoge@gmail.com",
			userSpaceName:  "Hoge family",
			password:       "password",
			validationKey:  cerrors.ValidationKey_AlreadyTaken,
			validationName: "email",
			beforeRun: func() {
				reg.NewUserRepository().Create(repository.UserCreateDTO{
					ID:            "hoge",
					AccountStatus: string(value.UserAccountStatus_Invited),
					Name:          "Hoge",
					Email:         "hoge@gmail.com",
					PasswordHash:  "hashed",
				})
			},
		},
		{
			name:           "Hoge",
			email:          "hoge@gmail.com",
			userSpaceName:  "Hoge family",
			password:       "",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "password",
		},
		{
			name:           "Hoge",
			email:          "hoge@gmail.com",
			userSpaceName:  "",
			password:       "password",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "user_space_name",
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"name":            util.StrToNilIfEmpty(test.name),
				"email":           util.StrToNilIfEmpty(test.email),
				"user_space_name": util.StrToNilIfEmpty(test.userSpaceName),
				"password":        util.StrToNilIfEmpty(test.password),
			},
		)

		// -------------------- execution --------------------
		if test.beforeRun != nil {
			test.beforeRun()
		}
		accountH := NewAccount()
		_, _, err := accountH.Signup(c, reg)

		// -------------------- assertion --------------------
		validationErr, ok := err.(cerrors.Validation)
		assert.True(t, ok)

		assert.Equal(t, test.validationKey, validationErr.Key)
		assert.Equal(t, test.validationName, validationErr.Name)
	}
}

func TestAccountSignupConfirm_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct{ nop string }{
		{},
	}

	for range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()

		accountUc, err := usecase.NewAccount(reg)
		assert.NoError(t, err)

		ret, err := accountUc.Signup(usecase.AccountSignupDTO{
			Name:          util.StrToPointer("Hoge Tarou"),
			Email:         util.StrToPointer("hoge@gmail.com"),
			UserSpaceName: util.StrToPointer("Hoge family"),
			Password:      util.StrToPointer("password"),
			SkipEmail:     util.BoolToPointer(true),
		})
		assert.NoError(t, err)

		ui, err := reg.NewUserInvitationRepository().FindOne(&repository.FindOption{
			Filters: []*repository.FindOptionFilter{
				{Query: "user_id = ?", Value: ret.UserID},
				{Query: "user_space_id = ?", Value: ret.UserSpaceID},
			},
		})
		assert.NoError(t, err)

		c := newGinContext(
			http.MethodPost,
			"/?id="+ui.ID,
		)

		// -------------------- execution --------------------
		accountH := NewAccount()
		_, _, err = accountH.SignupConfirm(c, reg)

		// -------------------- assertion --------------------
		assert.NoError(t, err)

		// user status changed
		user, err := reg.NewUserRepository().FindOne(&repository.FindOption{
			Filters: []*repository.FindOptionFilter{
				{Query: "id = ?", Value: ret.UserID},
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, value.UserAccountStatus_Confirmed, user.AccountStatus)
	}
}

func TestAccountSignupConfirm_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		validationKey  cerrors.ValidationKey
		validationName string
	}{
		{validationKey: cerrors.ValidationKey_Required, validationName: "id"},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()

		c := newGinContext(
			http.MethodGet,
			"/",
		)

		// -------------------- execution --------------------
		accountH := NewAccount()
		_, _, err := accountH.SignupConfirm(c, reg)

		// -------------------- assertion --------------------
		validationErr, ok := err.(cerrors.Validation)
		assert.True(t, ok)

		assert.Equal(t, test.validationKey, validationErr.Key)
		assert.Equal(t, test.validationName, validationErr.Name)
	}
}

func TestAccountLogin_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct{ nop string }{
		{},
	}

	for range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		uenv := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"email":    util.StrToPointer(uenv.User.Email),
				"password": util.StrToPointer("password"),
			},
		)

		// -------------------- execution --------------------
		accountH := NewAccount()
		status, data, err := accountH.Login(c, reg)

		// -------------------- assertion --------------------
		assert.Equal(t, http.StatusOK, status)
		assert.NoError(t, err)

		decodedRes := AccountLoginRes{}
		verifyJSONEncoding(data, &decodedRes)

		authUc := usecase.NewAuth(reg)
		ret, err := authUc.VerifyJWT(decodedRes.Token)
		assert.NoError(t, err)
		assert.Equal(t, uenv.User.ID, ret.UserID)
		assert.Equal(t, uenv.UserSpace.ID, ret.UserSpaceID)
	}
}

func TestAccountLogin_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		email          string
		password       string
		validationKey  cerrors.ValidationKey
		validationName string
	}{
		{
			email:          "",
			password:       "password",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "email",
		},
		{
			email:          "unknown@gmail.com",
			password:       "password",
			validationKey:  cerrors.ValidationKey_ResourceNotFound,
			validationName: "email",
		},
		{
			email:          "hoge@gmail.com",
			password:       "",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "password",
		},
		{
			email:          "hoge@gmail.com",
			password:       "wrong",
			validationKey:  cerrors.ValidationKey_Invalid,
			validationName: "password",
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		_ = api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodGet,
			"/",
			map[string]any{
				"email":    util.StrToNilIfEmpty(test.email),
				"password": util.StrToNilIfEmpty(test.password),
			},
		)

		// -------------------- execution --------------------
		accountH := NewAccount()
		_, _, err := accountH.Login(c, reg)

		// -------------------- assertion --------------------
		validationErr, ok := err.(cerrors.Validation)
		assert.True(t, ok)

		assert.Equal(t, test.validationKey, validationErr.Key)
		assert.Equal(t, test.validationName, validationErr.Name)
	}
}

func TestAccountInviteUser_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct{ email string }{
		{email: "hoge2@gmail.com"},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		uenv := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"email": util.StrToPointer(test.email),
			},
		)
		uenv.SetupAuthorization(c)

		// setup assertion
		api.MockMailer.
			EXPECT().
			Send(
				gomock.Any(),
				gomock.Cond(func(dto any) bool {
					return (dto).(interfaces.MailerSendDTO).To[0] == test.email
				}),
			)

		// -------------------- execution --------------------
		accountH := NewAccount()
		_, _, err := accountH.InviteUser(c, reg)

		// -------------------- assertion --------------------
		user, err := reg.NewUserRepository().FindOne(&repository.FindOption{
			Filters: []*repository.FindOptionFilter{
				{Query: "email = ?", Value: test.email},
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, value.UserAccountStatus_Invited, user.AccountStatus)

		ui, err := reg.NewUserInvitationRepository().FindOne(&repository.FindOption{
			Filters: []*repository.FindOptionFilter{
				{Query: "user_id = ?", Value: user.ID},
				{Query: "user_space_id = ?", Value: uenv.UserSpace.ID},
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, value.UserInvitationType_Invite, ui.Type)
	}
}

func TestAccountInviteUser_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		email          string
		validationKey  cerrors.ValidationKey
		validationName string
	}{
		{
			email:          "",
			validationKey:  cerrors.ValidationKey_Required,
			validationName: "email",
		},
		{
			email:          "invalid",
			validationKey:  cerrors.ValidationKey_InvalidFormat,
			validationName: "email",
		},
		{
			// Use the operating user's email
			email:          "hoge@gmail.com",
			validationKey:  cerrors.ValidationKey_AlreadyTaken,
			validationName: "email",
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		uenv := api.InstallBaseUserEnv()

		c := newGinContextWithBody(
			http.MethodGet,
			"/",
			map[string]any{
				"email": util.StrToNilIfEmpty(test.email),
			},
		)
		uenv.SetupAuthorization(c)

		// -------------------- execution --------------------
		accountH := NewAccount()
		_, _, err := accountH.InviteUser(c, reg)

		// -------------------- assertion --------------------
		validationErr, ok := err.(cerrors.Validation)
		assert.True(t, ok)

		assert.Equal(t, test.validationKey, validationErr.Key)
		assert.Equal(t, test.validationName, validationErr.Name)
	}
}

func TestAccountInviteUserConfirm_S(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct{ nop string }{
		{},
	}

	for range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		uenv := api.InstallBaseUserEnv()

		accountUc, err := usecase.NewAccount(reg)
		assert.NoError(t, err)

		inviteUserRet, err := accountUc.InviteUser(usecase.AccountInviteUserDTO{
			UserSpaceID: uenv.UserSpace.ID,
			Email:       util.StrToPointer("hoge2@gmail.com"),
			SkipEmail:   util.BoolToPointer(true),
		})
		assert.NoError(t, err)

		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			map[string]any{
				"name":          util.StrToPointer("Hoge Tarou 2"),
				"password":      util.StrToPointer("password"),
				"invitation_id": util.StrToPointer(inviteUserRet.InvitationID),
			},
		)

		// -------------------- execution --------------------
		accountH := NewAccount()
		_, data, err := accountH.InviteUserConfirm(c, reg)

		// -------------------- assertion --------------------
		assert.NoError(t, err)

		decodedRes := AccountInviteUserConfirmRes{}
		verifyJSONEncoding(data, &decodedRes)

		authUc := usecase.NewAuth(reg)
		ret, err := authUc.VerifyJWT(decodedRes.Token)

		// user status changed
		user, err := reg.NewUserRepository().FindOne(&repository.FindOption{
			Filters: []*repository.FindOptionFilter{
				{Query: "id = ?", Value: ret.UserID},
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, value.UserAccountStatus_Confirmed, user.AccountStatus)
	}
}

func TestAccountInviteUserConfirm_F_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	reg, api, err := testutil.SetupTestEnvironment(ctrl)
	assert.NoError(t, err)

	tests := []struct {
		name                   string
		password               string
		setInvitationID        bool
		useUnknownInvitationID bool
		validationKey          cerrors.ValidationKey
		validationName         string
	}{
		{
			name:                   "",
			password:               "password",
			setInvitationID:        true,
			useUnknownInvitationID: false,
			validationKey:          cerrors.ValidationKey_Required,
			validationName:         "name",
		},
		{
			name:                   "Hoge Tarou 2",
			password:               "",
			setInvitationID:        true,
			useUnknownInvitationID: false,
			validationKey:          cerrors.ValidationKey_Required,
			validationName:         "password",
		},
		{
			name:                   "Hoge Tarou 2",
			password:               "password",
			setInvitationID:        false,
			useUnknownInvitationID: false,
			validationKey:          cerrors.ValidationKey_Required,
			validationName:         "invitation-id",
		},
		{
			name:                   "Hoge Tarou 2",
			password:               "password",
			setInvitationID:        true,
			useUnknownInvitationID: true,
			validationKey:          cerrors.ValidationKey_ResourceNotFound,
			validationName:         "invitation-id",
		},
	}

	for _, test := range tests {
		// -------------------- preparation --------------------
		api.CleanupDB()
		uenv := api.InstallBaseUserEnv()

		accountUc, err := usecase.NewAccount(reg)
		assert.NoError(t, err)

		inviteUserRet, err := accountUc.InviteUser(usecase.AccountInviteUserDTO{
			UserSpaceID: uenv.UserSpace.ID,
			Email:       util.StrToPointer("hoge2@gmail.com"),
			SkipEmail:   util.BoolToPointer(true),
		})
		assert.NoError(t, err)

		body := map[string]any{
			"name":          util.StrToNilIfEmpty(test.name),
			"password":      util.StrToNilIfEmpty(test.password),
			"invitation_id": util.StrToNilIfEmpty(inviteUserRet.InvitationID),
		}
		if !test.setInvitationID {
			body["invitation_id"] = nil
		} else if test.useUnknownInvitationID {
			body["invitation_id"] = util.StrToPointer("invalid-id")
		}
		c := newGinContextWithBody(
			http.MethodPost,
			"/",
			body,
		)

		// -------------------- execution --------------------
		accountH := NewAccount()
		_, _, err = accountH.InviteUserConfirm(c, reg)

		// -------------------- assertion --------------------
		validationErr, ok := err.(cerrors.Validation)
		assert.True(t, ok)

		assert.Equal(t, test.validationKey, validationErr.Key)
		assert.Equal(t, test.validationName, validationErr.Name)
	}
}
