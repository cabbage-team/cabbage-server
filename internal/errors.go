package internal

import (
	"errors"
	"fmt"
	"net/http"
)

type ServerError struct {
	Status   int               // http 状态码
	Code     int               // 错误码
	Msg      string            // 错误信息
	ErrorMsg map[string]string //详细错误
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("server error: status=%d code=%d, msg=%s", e.Status, e.Code, e.Msg)
}

//	func (e *ServerError) Self() *ServerError {
//		return e
//	}
func (e *ServerError) Unwrap() error {
	return errors.New(e.Error())
}

func New(status int, code int, message string) *ServerError {
	return &ServerError{
		Code:     code,
		Status:   status,
		Msg:      message,
		ErrorMsg: map[string]string{},
	}
}

// 内置错误

// the record not found
var RecordNotFoundError = New(http.StatusNotFound, 0, "not found")

// 
var RecordUpdateError = New(http.StatusBadRequest, 0, "record update has wrong")

// 
var RequestParamsNotValidError = New(http.StatusBadRequest, 2, "the parameter is invalid")

// 
var UserNotFoundError = New(http.StatusNotFound, 3, "user not found")

// 
var UserNotAuthorized = New(http.StatusUnauthorized, 4, "用户身份未验证")

// 
var UserExprieError = New(http.StatusUnauthorized, 5, "用户身份过期")

// 
var UserRegisterError = New(http.StatusBadRequest, 8, "用户注册失败")

// 
var UserAlreadyExists = New(http.StatusBadRequest, 12, "用户已经存在")

// 
var UserLoginFail = New(http.StatusBadRequest,33,"用户登录失败")

// 
var UserNotExistsErr = New(http.StatusBadRequest,44,"用户不存在")

// 
var UserPasswordNotCorrectErr = New(http.StatusBadRequest,55,"用户密码错误")

// 
var LimitPermissionError = New(http.StatusForbidden, 6, "权限受限,无法操作")

// 
var InernalError = New(http.StatusInternalServerError, 7, "服务器内部错误")
// 
var CommentError = New(http.StatusBadRequest, 10, "comment error")
// 
var CommentNotFound = New(http.StatusNotFound, 10, "comment not found")
// 
var CommentOperatorError = New(http.StatusBadRequest, 11, "评论操作失败")
