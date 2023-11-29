package vilidate

import (
	"reflect"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

/**
 * 验证器 验证结构体值是否符合格式
 * @return {map[string]string} 当有未通过的字段 返回对应字段名称以及错误提示信息
 */
func Validators(target any) map[string]string {
	validater := validator.New()
	Type := reflect.TypeOf(target).Elem()    // 获取指向结构体的指针
	Values := reflect.ValueOf(target).Elem() // 获取指向结构体的指针
	errMsg := make(map[string]string)
	for i := 0; i < Type.NumField(); i++ {
		filed := Type.Field(i)
		name := filed.Name                       // 字段名称
		FieldValue := Values.FieldByName(name)   // 字段对象
		validateTag := filed.Tag.Get("validate") // tag值
		ErrorMsgTag := filed.Tag.Get("validateMsg")
		errs := validater.Var(FieldValue.Interface(), validateTag)
		if errs != nil {
			errMsg[strings.ToLower(filed.Name)] = ErrorMsgTag
		}
	}
	return errMsg
}