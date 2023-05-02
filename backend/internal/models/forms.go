package models

type FormLoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type FormRegisterUser struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,password"`
	PasswordRepeated string `json:"password_repeat" validate:"required,eqfield=Password"`
}
