package req

import (
	"github.com/go-playground/validator/v10"
)

type LoginValidator struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateReq struct {
	Username  string `json:"username" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
	Nickname  string `json:"nickname"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Roles     []int  `json:"roles"`
}

type UserResetPasswordReq struct {
	Password string `json:"password" binding:"required"`
}
type UserChangePasswordReq struct {
	Password string `json:"password" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}
func (r *LoginValidator) GetError(err error) string {
	errors := err.(validator.ValidationErrors)
	for _, e := range errors {
		switch e.Field() {
		case "Username":
			switch e.ActualTag() {
			case "required":
				return "the username field is required"
			case "lt":
				return "the len of username must less than 5"
			}
		case "Password":
			switch e.ActualTag() {
			case "required":
				return "the password field is required"
			}
		}
	}

	return "param is error"
}
