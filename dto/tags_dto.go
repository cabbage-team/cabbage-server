package dto

type TagDTO struct {
	// tag 名称
	Name string `json:"name" form:"name" validate:"required" validateMsg:"this field is required"`
	// tag 描述
	Description string `json:"description" form:"description" validate:"required" validateMsg:"this field is required"`
	// 创建者
	CreatedBy string `json:"createdby" form:"createdby" validate:"required" validateMsg:"this field is required"`
	// 封面 配图 图片连接地址
	Cover string `json:"cover" form:"cover" validate:"required" validateMsg:"this field is required"`
}