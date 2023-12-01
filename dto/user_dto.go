package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email" validateMsg:"the email is required"`
	Password string `json:"password" form:"password" validate:"required,min=10" validateMsg:"the password minimum 10 characters required"`
	Key      string `json:"key" form:"key" validate:"hexadecimal,required" validateMsg:"key must be a legal hexadecimal string"`
}

type SignupDTO struct {
	Name string `form:"name" json:"name" validate:"min=10,required" validateMsg:"the user name minimum 10 characters required"`
	Email    string `form:"email" json:"email" validate:"required,email" validateMsg:"the email is required"`
	Password string `json:"" form:"password" validate:"required,min=10" validateMsg:"the password minimum 10 characters required"`
}

type UserProfileDTO struct {
	Email string `json:"email" form:"email" validate:"email,required"`
}