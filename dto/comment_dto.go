package dto

type CommentOperatorDTO struct {
	Cid int `json:"cid" form:"cid" validate:"number,required"`
	Op int `json:"op" form:"op" validate:"oneof=1 2 4 8 16,required,number"`
}

type CommentDTO struct {
	Post int `json:"post"`
	Content string `json:"content"`
}