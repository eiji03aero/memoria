package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"memoria-api/registry"
	"memoria-api/usecase"
	"memoria-api/usecase/ccontext"
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
	authUc := usecase.NewAuth(reg)
	accountUc, err := usecase.NewAccount(reg)
	if err != nil {
		return
	}

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

type AccountLoginReq struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type AccountLoginRes struct {
	Token string `json:"token"`
}

func (h *Account) Login(c *gin.Context, reg registry.Registry) (status int, data any, err error) {
	accountUc, err := usecase.NewAccount(reg)
	if err != nil {
		return
	}

	authUc := usecase.NewAuth(reg)

	body := AccountLoginReq{}
	err = c.BindJSON(&body)
	if err != nil {
		return
	}

	ret, err := accountUc.Login(usecase.AccountLoginDTO{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return
	}

	jwt, err := authUc.CreateJWT(usecase.AuthCreateJWTDTO{
		UserID:      ret.UserID,
		UserSpaceID: ret.UserSpaceID,
	})
	if err != nil {
		return
	}

	res := AccountLoginRes{Token: jwt}
	return http.StatusOK, res, nil
}

type AccountInviteUserReq struct {
	Email *string `json:"email"`
}

func (h *Account) InviteUser(c *gin.Context, reg registry.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)
	accountUc, err := usecase.NewAccount(reg)
	if err != nil {
		return
	}

	body := AccountInviteUserReq{}
	err = c.BindJSON(&body)
	if err != nil {
		return
	}

	err = accountUc.InviteUser(usecase.AccountInviteUserDTO{
		Email:       body.Email,
		UserSpaceID: cctx.GetUserSpaceID(),
	})
	if err != nil {
		return
	}

	return
}

type AccountInviteUserConfirmReq struct {
	Name         *string `json:"name"`
	Password     *string `json:"password"`
	InvitationID *string `json:"invitation_id"`
}

type AccountInviteUserConfirmRes struct {
	Token string `json:"token"`
}

func (h *Account) InviteUserConfirm(c *gin.Context, reg registry.Registry) (status int, data any, err error) {
	authUc := usecase.NewAuth(reg)
	accountUc, err := usecase.NewAccount(reg)
	if err != nil {
		return
	}

	body := AccountInviteUserConfirmReq{}
	err = c.BindJSON(&body)
	if err != nil {
		return
	}

	ret, err := accountUc.InviteUserConfirm(usecase.AccountInviteUserConfirmDTO{
		Name:         body.Name,
		Password:     body.Password,
		InvitationID: body.InvitationID,
	})
	if err != nil {
		return
	}

	jwt, err := authUc.CreateJWT(usecase.AuthCreateJWTDTO{
		UserID:      ret.UserID,
		UserSpaceID: ret.UserSpaceID,
	})
	if err != nil {
		return
	}

	res := AccountLoginRes{Token: jwt}
	return http.StatusOK, res, nil
}
