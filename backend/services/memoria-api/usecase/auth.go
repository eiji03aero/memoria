package usecase

import (
	"time"

	"memoria-api/config"
	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/service"
	"memoria-api/registry"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(config.JWTSecretKey)

type Auth interface {
	CreateJWT(dto AuthCreateJWTDTO) (string, error)
	VerifyJWT(dto AuthVerifyJWTDTO) (userID string, userSpaceID string, err error)
}

type auth struct {
	userRepo repository.User
	userSvc  *service.User
}

func NewAuth(reg registry.Registry) Auth {
	return &auth{
		userRepo: reg.NewUserRepository(),
		userSvc:  reg.NewUserService(),
	}
}

type AuthCreateJWTDTO struct {
	UserID      string
	UserSpaceID string
}

func (u *auth) CreateJWT(dto AuthCreateJWTDTO) (tokenString string, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":      dto.UserID,
			"userSpaceID": dto.UserSpaceID,
			"exp":         time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
	)

	tokenString, err = token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type AuthVerifyJWTDTO struct {
	TokenString string
}

func (u *auth) VerifyJWT(dto AuthVerifyJWTDTO) (userID string, userSpaceID string, err error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		dto.TokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		},
	)
	if err != nil {
		err = cerrors.NewInternal(err.Error())
		return
	}

	if !token.Valid {
		err = cerrors.NewUnauthorized()
		return
	}

	userID = claims["userID"].(string)
	userSpaceID = claims["userSpaceID"].(string)

	return
}
