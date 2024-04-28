package usecase

import (
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/repository"
	"memoria-api/domain/service"
	"memoria-api/registry"
)

type Account interface {
	Signup(dto AccountSignupDTO) (userID string, userSpaceId string, err error)
	AddUserToUserSpace(dto AccountAddUserToUserSpaceDTO) error
}

type account struct {
	registry                  registry.Registry
	userRepo                  repository.User
	userSpaceRepo             repository.UserSpace
	userUserSpaceRelationRepo repository.UserUserSpaceRelation
}

func NewAccount(reg registry.Registry) Account {
	return &account{
		registry:                  reg,
		userRepo:                  reg.NewUserRepository(),
		userSpaceRepo:             reg.NewUserSpaceRepository(),
		userUserSpaceRelationRepo: reg.NewUserUserSpaceRelationRepository(),
	}
}

type AccountSignupDTO struct {
	Name          *string
	Email         *string
	Password      *string
	UserSpaceName *string
}

func (u *account) Signup(dto AccountSignupDTO) (userID string, userSpaceID string, err error) {
	u.registry.BeginTx()

	// name check
	if dto.Name == nil {
		err = cerrors.NewValidation("name is required")
		return
	}

	// email check
	if dto.Email == nil {
		err = cerrors.NewValidation("email is required")
		return
	}
	if _, err = mail.ParseAddress(*dto.Email); err != nil {
		err = cerrors.NewValidation("email format is invalid")
		return
	}

	// password check
	if dto.Password == nil {
		err = cerrors.NewValidation("password is required")
		return
	}

	// user space name check
	if dto.UserSpaceName == nil {
		err = cerrors.NewValidation("user_space_name is required")
		return
	}

	// generate id for user
	userID = service.GenerateUlid()

	// assuming password sent was pretty strong
	hashed, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), 10)
	if err != nil {
		return
	}

	// create user
	err = u.userRepo.Create(repository.UserCreateDTO{
		ID:           userID,
		Name:         *dto.Name,
		Email:        *dto.Email,
		PasswordHash: string(hashed[:]),
	})
	if err != nil {
		err = cerrors.NewInternal(fmt.Sprintf("failed to create user: %s", err.Error()))
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
