package models

type User struct{
	ID uint
	Name string `validate:"required"`
	Email string `validate:"required,email"`
	Password string `validate:"required"`
}

type RegisterUserInput struct{
	User
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
