package dto

type CommentOperatorDTO struct {
	Cid int `json:"cid" form:"cid" validate:"number,required"`
	Op int `json:"op" form:"op" validate:"oneof=1 2 4 8 16,required,number"`
}

type CommentDTO struct {
	Post int `json:"post" validate:"number,required" validateMsg:"this field must be number"` 
	Content string `json:"content" validate:"min=1,required" validateMsg:"this field must be more than 1 character"`
}