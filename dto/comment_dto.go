package dto

type CommentOperatorDTO struct {
	Op int `json:"op" form:"op" validate:"oneof=1 2 4 8 16,required,number"`
}