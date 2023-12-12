package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email" validateMsg:"the email is required a legal email"`
	Password string `json:"password" form:"password" validate:"required,min=10" validateMsg:"the password minimum 10 characters required"`
	Key      string `json:"key" form:"key" validate:"hexadecimal,required" validateMsg:"key must be a legal hexadecimal string"`
}

type SignupDTO struct {
	Name string `form:"name" json:"name" validate:"min=4,required" validateMsg:"the user name minimum 10 characters required"`
	Email    string `form:"email" json:"email" validate:"required,email" validateMsg:"the email is required a legal email"`
	Password string `json:"" form:"password" validate:"required,min=10" validateMsg:"the password minimum 10 characters required"`
}

type UserProfileDTO struct {
	Email string `json:"email" form:"email" validate:"email,required" validateMsg:"this must be a valid email address"`
}

type NickNameDTO struct {
	Name string `json:"name" form:"name" validate:"min=4,required" validateMsg:"the user name minimum 10 characters required"`
}

type FollowListDTO struct {
	Page int `json:"page" form:"page" validate:"min=1" validateMsg:"the minimum number of pages is 1"`
	Size int `json:"size" form:"size" validate:"min=1,max=20" validateMsg:"the size out of range, must between 1 and 20"`
	UID int64 `json:"uid" form:"uid" validate:"min=1" validateMsg:"the minimum number of uid is 1"`
}