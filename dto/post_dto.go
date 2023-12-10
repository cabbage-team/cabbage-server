package dto

type PostDTO struct {
	Title   string   `json:"title" form:"title" validate:"required,min=8" validateMsg:"this field must be more than 8 character"`
	Content string   `json:"Content" form:"Content" validate:"required,min=30" validateMsg:"this field must be more than 30 character"`
	Tags    []string `json:"Tags" form:"Tags" validate:"dive,omitempty" validateMsg:"this field must be a string array"`
}
