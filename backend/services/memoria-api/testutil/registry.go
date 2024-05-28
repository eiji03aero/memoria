package testutil

import (
	"memoria-api/application/ccontext"
	"memoria-api/application/usecase"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/registry"
	"memoria-api/testutil/mock"
	"memoria-api/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func SetupTestEnvironment(ctrl *gomock.Controller) (reg interfaces.Registry, api *TestEnvAPI, err error) {
	// -------------------- build real registry --------------------
	realReg, err := registry.NewBuilder().Build()
	if err != nil {
		return
	}

	db := realReg.(*registry.Registry).DB

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
	mockReg.EXPECT().NewAlbumRepository().DoAndReturn(realReg.NewAlbumRepository).AnyTimes()
	mockReg.EXPECT().NewUserSpaceAlbumRelationRepository().DoAndReturn(realReg.NewUserSpaceAlbumRelationRepository).AnyTimes()
	mockReg.EXPECT().NewMediumRepository().DoAndReturn(realReg.NewMediumRepository).AnyTimes()
	mockReg.EXPECT().NewAlbumMediumRelationRepository().DoAndReturn(realReg.NewAlbumMediumRelationRepository).AnyTimes()
	mockReg.EXPECT().NewUserSpaceActivityRepository().DoAndReturn(realReg.NewUserSpaceActivityRepository).AnyTimes()
	// service
	mockReg.EXPECT().NewUserService().DoAndReturn(realReg.NewUserService).AnyTimes()
	mockReg.EXPECT().NewUserSpaceService().DoAndReturn(realReg.NewUserSpaceService).AnyTimes()
	mockReg.EXPECT().NewUserUserSpaceRelationService().DoAndReturn(realReg.NewUserUserSpaceRelationService).AnyTimes()
	mockReg.EXPECT().NewUserInvitationService().DoAndReturn(realReg.NewUserInvitationService).AnyTimes()
	mockReg.EXPECT().NewAlbumService().DoAndReturn(realReg.NewAlbumService).AnyTimes()
	mockReg.EXPECT().NewMediumService().DoAndReturn(realReg.NewMediumService).AnyTimes()
	mockReg.EXPECT().NewUserSpaceActivityService().DoAndReturn(realReg.NewUserSpaceActivityService).AnyTimes()

	// mocked
	mockMailer := mock.NewMockMailer(ctrl)
	mockReg.EXPECT().NewSESMailer().Return(mockMailer, nil).AnyTimes()
	mockS3Client := mock.NewMockS3Client(ctrl)
	mockReg.EXPECT().NewS3Client().Return(mockS3Client).AnyTimes()
	mockBGJobInvoker := mock.NewMockBGJobInvoker(ctrl)
	mockReg.EXPECT().NewBGJobInvoker().Return(mockBGJobInvoker).AnyTimes()

	reg = mockReg

	// -------------------- api --------------------
	api = &TestEnvAPI{
		db:               db,
		registry:         reg,
		MockMailer:       mockMailer,
		MockS3Client:     mockS3Client,
		MockBGJobInvoker: mockBGJobInvoker,
	}

	return
}

type TestEnvAPI struct {
	db               *gorm.DB
	registry         interfaces.Registry
	MockMailer       *mock.MockMailer
	MockS3Client     *mock.MockS3Client
	MockBGJobInvoker *mock.MockBGJobInvoker
}

func (t *TestEnvAPI) DB() *gorm.DB {
	return t.db
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
