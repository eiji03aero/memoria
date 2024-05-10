package testutil

import (
	"context"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/caws"
	"memoria-api/infra/db"
	"memoria-api/registry"
	"memoria-api/testutil/mock"
	"memoria-api/usecase"
	"memoria-api/usecase/ccontext"
	"memoria-api/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func SetupTestEnvironment(ctrl *gomock.Controller) (reg registry.Registry, api *TestEnvAPI, err error) {
	ctx := context.Background()

	// -------------------- build real registry --------------------
	Db, err := db.New()
	if err != nil {
		return
	}

	awsCfg, err := caws.LoadConfig(ctx)
	if err != nil {
		return
	}

	realReg := registry.NewRegistry(registry.NewRegistryDTO{
		DB:     Db,
		AWSCfg: awsCfg,
	})

	// -------------------- build mock registry --------------------
	mockReg := mock.NewMockRegistry(ctrl)

	// database
	mockReg.EXPECT().BeginTx().DoAndReturn(realReg.BeginTx).AnyTimes()
	mockReg.EXPECT().RollbackTx().DoAndReturn(realReg.RollbackTx).AnyTimes()
	mockReg.EXPECT().CommitTx().DoAndReturn(realReg.CommitTx).AnyTimes()
	mockReg.EXPECT().CloseDB().DoAndReturn(realReg.CloseDB).AnyTimes()
	// repository
	mockReg.EXPECT().NewUserRepository().DoAndReturn(realReg.NewUserRepository).AnyTimes()
	mockReg.EXPECT().NewUserSpaceRepository().DoAndReturn(realReg.NewUserSpaceRepository).AnyTimes()
	mockReg.EXPECT().NewUserUserSpaceRelationRepository().DoAndReturn(realReg.NewUserUserSpaceRelationRepository).AnyTimes()
	mockReg.EXPECT().NewUserInvitationRepository().DoAndReturn(realReg.NewUserInvitationRepository).AnyTimes()
	// service
	mockReg.EXPECT().NewUserService().DoAndReturn(realReg.NewUserService).AnyTimes()
	mockReg.EXPECT().NewUserSpaceService().DoAndReturn(realReg.NewUserSpaceService).AnyTimes()
	mockReg.EXPECT().NewUserUserSpaceRelationService().DoAndReturn(realReg.NewUserUserSpaceRelationService).AnyTimes()
	mockReg.EXPECT().NewUserInvitationService().DoAndReturn(realReg.NewUserInvitationService).AnyTimes()

	// mocked

	mockMailer := mock.NewMockMailer(ctrl)
	mockReg.EXPECT().NewSESMailer().Return(mockMailer, nil).AnyTimes()

	reg = mockReg

	// -------------------- api --------------------
	api = &TestEnvAPI{
		db:         Db,
		registry:   reg,
		MockMailer: mockMailer,
	}

	return
}

type TestEnvAPI struct {
	db         *gorm.DB
	registry   registry.Registry
	MockMailer *mock.MockMailer
}

func (t *TestEnvAPI) CleanupDB() {
	CleanupDB(t.db)
}

type TestEnvAPIAddUserDTO struct {
	Name        string
	Email       string
	Password    string
	UserSpaceID string
}

type TestEnvAPIAddUserRet struct {
	User *model.User
}

func (t *TestEnvAPI) AddUser(dto TestEnvAPIAddUserDTO) (ret TestEnvAPIAddUserRet) {
	accountUc, err := usecase.NewAccount(t.registry)
	if err != nil {
		panic(err)
	}

	inviteUserRet, err := accountUc.InviteUser(usecase.AccountInviteUserDTO{
		UserSpaceID: dto.UserSpaceID,
		Email:       util.StrToPointer(dto.Email),
		SkipEmail:   util.BoolToPointer(true),
	})
	if err != nil {
		return
	}

	inviteUserConfirmRet, err := accountUc.InviteUserConfirm(usecase.AccountInviteUserConfirmDTO{
		Name:         util.StrToPointer(dto.Name),
		Password:     util.StrToPointer(dto.Password),
		InvitationID: util.StrToPointer(inviteUserRet.InvitationID),
	})
	if err != nil {
		return
	}

	user, err := t.registry.NewUserRepository().FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: inviteUserConfirmRet.UserID},
		},
	})
	if err != nil {
		return
	}

	ret.User = user
	return
}

type UserEnv struct {
	User      *model.User
	UserSpace *model.UserSpace
	JWT       string
}

func (t *TestEnvAPI) InstallBaseUserEnv() (userEnv UserEnv) {
	userRepo := t.registry.NewUserRepository()
	userSpaceRepo := t.registry.NewUserSpaceRepository()
	accountUc, err := usecase.NewAccount(t.registry)
	if err != nil {
		panic(err)
	}
	authUc := usecase.NewAuth(t.registry)
	if err != nil {
		panic(err)
	}

	signupRet, err := accountUc.Signup(usecase.AccountSignupDTO{
		Name:          util.StrToPointer("Hoge Tarou"),
		Email:         util.StrToPointer("hoge@gmail.com"),
		UserSpaceName: util.StrToPointer("Hoge family"),
		Password:      util.StrToPointer("password"),
		SkipEmail:     util.BoolToPointer(true),
	})
	if err != nil {
		panic(err)
	}

	_, err = accountUc.SignupConfirm(usecase.AccountSignupConfirmDTO{
		ID: util.StrToPointer(signupRet.InvitationID),
	})
	if err != nil {
		panic(err)
	}

	user, err := userRepo.FindByID(signupRet.UserID)
	if err != nil {
		panic(err)
	}

	userSpace, err := userSpaceRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: signupRet.UserSpaceID},
		},
	})
	if err != nil {
		panic(err)
	}

	jwt, err := authUc.CreateJWT(usecase.AuthCreateJWTDTO{
		UserID:      user.ID,
		UserSpaceID: userSpace.ID,
	})
	if err != nil {
		panic(err)
	}

	userEnv.User = user
	userEnv.UserSpace = userSpace
	userEnv.JWT = jwt
	return
}

func (u *UserEnv) SetupAuthorization(c *gin.Context) {
	c.Request.Header.Add("Authorization", "Bearer "+u.JWT)
	ccontext.SetUserID(c, u.User.ID)
	ccontext.SetUserSpaceID(c, u.UserSpace.ID)
}
