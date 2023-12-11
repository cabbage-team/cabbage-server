package dto

type AddPlatformDTO struct {
	Name []string `json:"name" form:"name" validate:"dive,required,min=1"`
}

type StatMonthDataDTO struct {
	Month int `json:"month" form:"month" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10 11 12" validateMsg:"this field must be a valid month number"`
}
