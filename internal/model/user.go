package model

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type (
	UserPaylod struct {
		Username    string `json:"username" validate:"required,min=5,max=255"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number" validate:"required,max=12,numeric"`
		Address     string `json:"address" validate:"required"`
		DoB         string `json:"dob" validate:"required"`
		Password    string `json:"password" validate:"required,min=7,max=255"`
		Fullname    string `json:"fullname" validate:"required,max=255"`
		Role        string
	}

	LoginPayload struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=7,max=255"`
	}
)

func (u UserPaylod) Validate() error {
	return Validate.Struct(u)
}

func (u LoginPayload) Validate() error {
	return Validate.Struct(u)
}

type (
	UserResponse struct {
		ID          int32  `json:"id"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		DoB         string `json:"dob"`
		Password    string `json:"-"`
		Fullname    string `json:"fullname"`
		Role        int32  `json:"Role"`
	}

	LoginResponse struct {
		ActiveToken  string `json:"active_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
