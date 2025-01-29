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
)

func (u UserPaylod) Validate() error {
	return Validate.Struct(u)
}
