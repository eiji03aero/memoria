package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"memoria-api/registry"
	"memoria-api/usecase"
)

type Account struct{}

func NewAccount() *Account {
	return &Account{}
}

type AccountSignupReq struct {
	Name          *string `json:"name"`
	Email         *string `json:"email"`
	Password      *string `json:"password"`
	UserSpaceName *string `json:"user_space_name"`
}

type AccountSignupRes struct {
	Token string `json:"token"`
}

func (h *Account) Signup(c *gin.Context, reg registry.Registry) (status int, data any, err error) {
	accountUc, err := usecase.NewAccount(reg)
	if err != nil {
		return
	}

	authUc := usecase.NewAuth(reg)

	body := AccountSignupReq{}
	err = c.BindJSON(&body)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	userID, userSpaceID, err := accountUc.Signup(usecase.AccountSignupDTO{
		Name:          body.Name,
		Email:         body.Email,
		Password:      body.Password,
		UserSpaceName: body.UserSpaceName,
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	jwt, err := authUc.CreateJWT(usecase.AuthCreateJWTDTO{
		UserID:      userID,
		UserSpaceID: userSpaceID,
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	res := AccountSignupRes{Token: jwt}
	return http.StatusOK, res, nil
}

type AccountSignupConfirmReq struct {
	ID *string `form:"id"`
}

func (h *Account) SignupConfirm(c *gin.Context, reg registry.Registry) (status int, data any, err error) {
	accountUc, err := usecase.NewAccount(reg)
	if err != nil {
		return
	}

	query := AccountSignupConfirmReq{}
	err = c.ShouldBindQuery(&query)
	if err != nil {
		return
	}

	ret, err := accountUc.SignupConfirm(usecase.AccountSignupConfirmDTO{
		ID: query.ID,
	})
	if err != nil {
		c.Redirect(http.StatusSeeOther, ret.RedirectURL)
		return
	}

	c.Redirect(http.StatusSeeOther, ret.RedirectURL)
	return
}
