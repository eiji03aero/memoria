package usecase

import (
	"context"
	"fmt"
	"log"
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"memoria-api/config"
	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/service"
	"memoria-api/domain/value"
	"memoria-api/registry"
)

type Account interface {
	Signup(dto AccountSignupDTO) (userID string, userSpaceId string, err error)
	SignupConfirm(dto AccountSignupConfirmDTO) (ret AccountSignupConfirmRet, err error)
	AddUserToUserSpace(dto AccountAddUserToUserSpaceDTO) error
}

type account struct {
	registry                  registry.Registry
	mailer                    interfaces.Mailer
	userRepo                  repository.User
	userSpaceRepo             repository.UserSpace
	userUserSpaceRelationRepo repository.UserUserSpaceRelation
	userInvitationRepo        repository.UserInvitation
	userSvc                   *service.User
}

func NewAccount(reg registry.Registry) (u Account, err error) {
	mailer, err := reg.NewSESMailer()
	if err != nil {
		return
	}
	u = &account{
		registry:                  reg,
		mailer:                    mailer,
		userRepo:                  reg.NewUserRepository(),
		userSpaceRepo:             reg.NewUserSpaceRepository(),
		userUserSpaceRelationRepo: reg.NewUserUserSpaceRelationRepository(),
		userInvitationRepo:        reg.NewUserInvitationRepository(),
		userSvc:                   reg.NewUserService(),
	}
	return
}

type AccountSignupDTO struct {
	Name          *string
	Email         *string
	Password      *string
	UserSpaceName *string
}

func (u *account) Signup(dto AccountSignupDTO) (userID string, userSpaceID string, err error) {
	ctx := context.Background()
	u.registry.BeginTx()

	// -------------------- validations --------------------
	if dto.Name == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "name",
		})
		return
	}

	if dto.Email == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "email",
		})
		return
	}
	_, err = mail.ParseAddress(*dto.Email)
	if err != nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_InvalidFormat,
			Name: "email",
		})
		return
	}
	isEmailTaken, err := u.userSvc.IsEmailTaken(service.UserIsEmailTakenDTO{Email: *dto.Email})
	if err != nil {
		return
	}
	if isEmailTaken {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_AlreadyTaken,
			Name: "email",
		})
		return
	}

	if dto.Password == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_InvalidFormat,
			Name: "password",
		})
		return
	}

	if dto.UserSpaceName == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_InvalidFormat,
			Name: "user_space_name",
		})
		return
	}

	// -------------------- execution --------------------
	// generate id for user
	userID = service.GenerateUlid()

	// assuming password sent was pretty strong
	hashed, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), 10)
	if err != nil {
		return
	}

	// create user
	err = u.userRepo.Create(repository.UserCreateDTO{
		ID:            userID,
		AccountStatus: string(value.UserAccountStatus_Invited),
		Name:          *dto.Name,
		Email:         *dto.Email,
		PasswordHash:  string(hashed[:]),
	})
	if err != nil {
		err = cerrors.NewInternal(fmt.Sprintf("failed to create user: %s", err.Error()))
		log.Println(err)
		log.Println(err.Error())
		u.registry.RollbackTx()
		return
	}

	// generate id for user space
	userSpaceID = service.GenerateUlid()

	// create user space
	err = u.userSpaceRepo.Create(repository.UserSpaceCreateDTO{
		ID:   userSpaceID,
		Name: *dto.UserSpaceName,
	})
	if err != nil {
		err = cerrors.NewInternal(fmt.Sprintf("failed to create user space: %s", err.Error()))
		u.registry.RollbackTx()
		return
	}

	// add user to the user space
	err = u.AddUserToUserSpace(AccountAddUserToUserSpaceDTO{
		UserID:      userID,
		UserSpaceID: userSpaceID,
	})
	if err != nil {
		err = cerrors.NewInternal(fmt.Sprintf("failed to add user to user space: %s", err.Error()))
		u.registry.RollbackTx()
		return
	}

	// create invitation
	invitationID := service.GenerateUlid()
	err = u.userInvitationRepo.Create(repository.UserInvitationCreateDTO{
		ID:     invitationID,
		UserID: userID,
		Type:   string(value.UserInvitationType_Signup),
	})

	// send confirm email
	confirmUrl := config.Host + fmt.Sprintf("/api/public/signup-confirm?id=%s", invitationID)
	body := fmt.Sprintf(`
Thanks for signing up memoria.<br/>

Please open the url below to complete signup process and start using memoria.<br/>
<a href="%s">%s</a>
`,
		confirmUrl, confirmUrl,
	)
	u.mailer.Send(ctx, interfaces.MailerSendDTO{
		From:    config.NoReplyEmailAddress,
		To:      []string{*dto.Email},
		Subject: "Memoria - Please confirm your email address",
		Body:    body,
	})

	return
}

type AccountSignupConfirmDTO struct {
	ID *string
}

type AccountSignupConfirmRet struct {
	RedirectURL string
}

func (u *account) SignupConfirm(dto AccountSignupConfirmDTO) (ret AccountSignupConfirmRet, err error) {
	setErrorURL := func() {
		ret.RedirectURL = config.ClientHost + "/internal-server-error"
	}

	if dto.ID == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "id",
		})
		setErrorURL()
		return
	}

	// confirm invitation exists
	userInvitation, err := u.userInvitationRepo.FindByID(repository.UserInvitationFindByIDDTO{
		ID: *dto.ID,
	})
	if err != nil {
		setErrorURL()
		return
	}

	user, err := u.userRepo.FindByID(repository.UserFindByIDDTO{
		ID: userInvitation.UserID,
	})
	if err != nil {
		setErrorURL()
		return
	}

	user.SetAccountStatus(value.UserAccountStatus_Confirmed)
	err = u.userRepo.Update(user)
	if err != nil {
		setErrorURL()
		return
	}

	ret.RedirectURL = config.ClientHost + "/signup-thanks"
	return
}

type AccountAddUserToUserSpaceDTO struct {
	UserID      string
	UserSpaceID string
}

func (u *account) AddUserToUserSpace(dto AccountAddUserToUserSpaceDTO) (err error) {
	// should we have validation over id?

	err = u.userUserSpaceRelationRepo.Create(repository.UserUserSpaceRelationCreateDTO{
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
	})

	return
}
