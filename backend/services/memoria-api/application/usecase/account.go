package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"memoria-api/config"
	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/service"
	"memoria-api/domain/value"
)

type Account interface {
	Signup(dto AccountSignupDTO) (ret AccountSignupRet, err error)
	SignupConfirm(dto AccountSignupConfirmDTO) (ret AccountSignupConfirmRet, err error)
	AddUserToUserSpace(dto AccountAddUserToUserSpaceDTO) error
	Login(dto AccountLoginDTO) (ret AccountLoginRet, err error)
	InviteUser(dto AccountInviteUserDTO) (ret AccountInviteUserRet, err error)
	InviteUserConfirm(dto AccountInviteUserConfirmDTO) (ret AccountInviteUserConfirmRet, err error)
}

type account struct {
	registry                  interfaces.Registry
	mailer                    interfaces.Mailer
	userRepo                  repository.User
	userSpaceRepo             repository.UserSpace
	userUserSpaceRelationRepo repository.UserUserSpaceRelation
	userInvitationRepo        repository.UserInvitation
	userSvc                   svc.User
	userInvitationSvc         svc.UserInvitation
	userSpaceSvc              svc.UserSpace
	userUserSpaceRelationSvc  svc.UserUserSpaceRelation
	usaSvc                    svc.UserSpaceActivity
}

func NewAccount(reg interfaces.Registry) (u Account, err error) {
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
		userInvitationSvc:         reg.NewUserInvitationService(),
		userSpaceSvc:              reg.NewUserSpaceService(),
		userUserSpaceRelationSvc:  reg.NewUserUserSpaceRelationService(),
		usaSvc:                    reg.NewUserSpaceActivityService(),
	}
	return
}

type AccountSignupDTO struct {
	Name          *string
	Email         *string
	Password      *string
	UserSpaceName *string
	SkipEmail     *bool
}

type AccountSignupRet struct {
	UserID       string
	UserSpaceID  string
	InvitationID string
}

func (u *account) Signup(dto AccountSignupDTO) (ret AccountSignupRet, err error) {
	ctx := context.Background()

	// -------------------- validation --------------------
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
	isEmailExists, err := u.userRepo.Exists(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "email = ?", Value: *dto.Email},
		},
	})
	if isEmailExists {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_AlreadyTaken,
			Name: "email",
		})
		return
	}
	if err != nil {
		return
	}

	if dto.Password == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "password",
		})
		return
	}

	if dto.UserSpaceName == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "user_space_name",
		})
		return
	}

	// -------------------- execution --------------------
	// generate id for user
	userID := service.GenerateUlid()

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
		return
	}

	// generate id for user space
	userSpaceID := service.GenerateUlid()

	// create user space
	err = u.userSpaceRepo.Create(repository.UserSpaceCreateDTO{
		ID:   userSpaceID,
		Name: *dto.UserSpaceName,
	})
	if err != nil {
		err = cerrors.NewInternal(fmt.Sprintf("failed to create user space: %s", err.Error()))
		return
	}

	// add user to the user space
	err = u.AddUserToUserSpace(AccountAddUserToUserSpaceDTO{
		UserID:      userID,
		UserSpaceID: userSpaceID,
	})
	if err != nil {
		err = cerrors.NewInternal(fmt.Sprintf("failed to add user to user space: %s", err.Error()))
		return
	}

	// create invitation
	invitationID := service.GenerateUlid()
	err = u.userInvitationRepo.Create(repository.UserInvitationCreateDTO{
		ID:          invitationID,
		UserID:      userID,
		UserSpaceID: userSpaceID,
		Type:        string(value.UserInvitationType_Signup),
	})

	shouldSendEmail := dto.SkipEmail == nil || !*dto.SkipEmail
	if shouldSendEmail {

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
	}

	ret.UserID = userID
	ret.UserSpaceID = userSpaceID
	ret.InvitationID = invitationID
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

	// -------------------- validation --------------------
	if dto.ID == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "id",
		})
		setErrorURL()
		return
	}

	// -------------------- execution --------------------
	userInvitation, err := u.userInvitationRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: *dto.ID},
		},
	})
	if err != nil {
		setErrorURL()
		return
	}

	user, err := u.userRepo.FindByID(userInvitation.UserID)
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

type AccountLoginDTO struct {
	Email    *string
	Password *string
}

type AccountLoginRet struct {
	UserID      string
	UserSpaceID string
}

func (u *account) Login(dto AccountLoginDTO) (ret AccountLoginRet, err error) {
	// -------------------- validations --------------------
	if dto.Email == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "email",
		})
		return
	}

	if dto.Password == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "password",
		})
		return
	}

	// -------------------- execution --------------------
	user, err := u.userRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "email = ?", Value: *dto.Email},
		},
	})
	if errors.As(err, &cerrors.ResourceNotFound{}) {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_ResourceNotFound,
			Name: "email",
		})
		return
	}
	if err != nil {
		return
	}

	uusr, err := u.userUserSpaceRelationRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "user_id = ?", Value: user.ID},
		},
	})
	if err != nil {
		return
	}

	matched, err := service.CompareHashAndPassword([]byte(user.PasswordHash), []byte(*dto.Password))
	if !matched {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Invalid,
			Name: "password",
		})
		return
	}
	if err != nil {
		return
	}

	ret.UserID = user.ID
	ret.UserSpaceID = uusr.UserSpaceID
	return
}

type AccountInviteUserDTO struct {
	UserSpaceID string
	Email       *string
	SkipEmail   *bool
}

type AccountInviteUserRet struct {
	InvitationID string
}

func (u *account) InviteUser(dto AccountInviteUserDTO) (ret AccountInviteUserRet, err error) {
	ctx := context.Background()

	// -------------------- validation --------------------
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

	userExists, err := u.userRepo.EmailExistsInUserSpace(dto.UserSpaceID, *dto.Email)
	if userExists {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_AlreadyTaken,
			Name: "email",
		})
		return
	}
	if err != nil {
		return
	}

	// -------------------- execution --------------------
	userID := service.GenerateUlid()

	err = u.userRepo.Create(repository.UserCreateDTO{
		ID:            userID,
		AccountStatus: string(value.UserAccountStatus_Invited),
		Name:          "Invited user",
		Email:         *dto.Email,
		PasswordHash:  "",
	})
	if err != nil {
		return
	}

	err = u.AddUserToUserSpace(AccountAddUserToUserSpaceDTO{
		UserID:      userID,
		UserSpaceID: dto.UserSpaceID,
	})
	if err != nil {
		err = cerrors.NewInternal(fmt.Sprintf("failed to add user to user space: %s", err.Error()))
		return
	}

	invitationID := service.GenerateUlid()
	err = u.userInvitationRepo.Create(repository.UserInvitationCreateDTO{
		ID:          invitationID,
		UserID:      userID,
		UserSpaceID: dto.UserSpaceID,
		Type:        string(value.UserInvitationType_Invite),
	})
	if err != nil {
		return
	}

	shouldSendEmail := dto.SkipEmail == nil || !*dto.SkipEmail
	if shouldSendEmail {
		confirmUrl := config.ClientHost + "/invite-user-confirm?id=" + invitationID
		body := fmt.Sprintf(`
Hello this is memoria.<br/>

You have been invited to a user space in memoria.<br />
Please open below url to join them! <br />
<a href="%s">%s</a>
`,
			confirmUrl, confirmUrl,
		)
		u.mailer.Send(ctx, interfaces.MailerSendDTO{
			From:    config.NoReplyEmailAddress,
			To:      []string{*dto.Email},
			Subject: "Memoria - Please confirm your invitation",
			Body:    body,
		})
	}

	ret.InvitationID = invitationID
	return
}

type AccountInviteUserConfirmDTO struct {
	Name         *string
	Password     *string
	InvitationID *string
}

type AccountInviteUserConfirmRet struct {
	UserID      string
	UserSpaceID string
}

func (u *account) InviteUserConfirm(dto AccountInviteUserConfirmDTO) (ret AccountInviteUserConfirmRet, err error) {
	// -------------------- validation --------------------
	if dto.Name == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "name",
		})
		return
	}

	if dto.Password == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "password",
		})
		return
	}

	if dto.InvitationID == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "invitation-id",
		})
		return
	}

	// -------------------- execution --------------------
	ui, err := u.userInvitationRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: *dto.InvitationID},
		},
	})
	if errors.As(err, &cerrors.ResourceNotFound{}) {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_ResourceNotFound,
			Name: "invitation-id",
		})
		return
	}
	if err != nil {
		return
	}

	user, err := u.userRepo.FindByID(ui.UserID)
	if err != nil {
		return
	}

	hashed, err := service.GenerateHashedPassword([]byte(*dto.Password))
	if err != nil {
		return
	}

	user.Name = *dto.Name
	user.PasswordHash = string(hashed)
	user.SetAccountStatus(value.UserAccountStatus_Confirmed)
	err = u.userRepo.Update(user)
	if err != nil {
		return
	}

	err = u.usaSvc.CreateInviteUserJoined(svc.UserSpaceActivityCreateInviteUserJoined{
		UserSpaceID: ui.UserSpaceID,
		UserID:      ui.UserID,
	})
	if err != nil {
		return
	}

	ret.UserID = ui.UserID
	ret.UserSpaceID = ui.UserSpaceID
	return
}
