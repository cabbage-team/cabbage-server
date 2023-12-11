package dto

type PageDTO struct {
	Page int `json:"page" form:"page" validate:"min=1" validateMsg:"the minimum number of pages is 1"`
	Size int `json:"size" form:"size" validate:"min=1,max=20" validateMsg:"the size out of range, must between 1 and 20"`
}