package usecase

import (
	"time"

	"memoria-api/config"
	"memoria-api/domain/cerrors"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(config.JWTSecretKey)

type Auth interface {
	CreateJWT(userID string, userSpaceID string) (string, error)
	VerifyJWT(tokenString string) (userID string, userSpaceID string, err error)
}

type auth struct{}

func NewAuth() Auth {
	return &auth{}
}

func (u *auth) CreateJWT(userID string, userSpaceID string) (tokenString string, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":      userID,
			"userSpaceID": userSpaceID,
			"exp":         time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
	)

	tokenString, err = token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *auth) VerifyJWT(tokenString string) (userID string, userSpaceID string, err error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
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
