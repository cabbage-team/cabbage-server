package internal

import (
	"cabbage-server/common/utils"
	"cabbage-server/i18n"
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
	return fmt.Sprintf(
		"server error: status=%d code=%d msg=%s",
		e.Status,
		e.Code,
		e.Msg,
	)
}

//	func (e *ServerError) Self() *ServerError {
//		return e
//	}
func (e *ServerError) Unwrap() error {
	return errors.New(e.Error(),
)
}

func New(status int,
 code int,
 message string) *ServerError {
	return &ServerError{
		Code:     code,
		Status:   status,
		Msg:      message,
		ErrorMsg: map[string]string{},
	}
}

// 内置错误

// the record not found
var RecordNotFoundError = New(
	http.StatusNotFound,
	1000,
	"not found",
)

var RecordUpdateError = New(
	http.StatusBadRequest,
	utils.ToInt(i18n.I18N.T("Status.RecordNotFound")),
	i18n.I18N.T("Error.RecordNotFound"),
)

var RequestParamsNotValidError = New(
	http.StatusBadRequest,
	utils.ToInt(i18n.I18N.T("Status.ParamsNotValid")),
	i18n.I18N.T("Error.ParamsNotValid"),
)

var UserNotFoundError = New(
	http.StatusNotFound,
	utils.ToInt(i18n.I18N.T("Status.UserNotFound")),
	i18n.I18N.T("Error.UserNotFound"),
)

var UserNotAuthorized = New(http.StatusUnauthorized,
	utils.ToInt(i18n.I18N.T("Status.UserNotAuthorized")),
	i18n.I18N.T("Error.UserNotAuthorized"),
)

var UserExprieError = New(http.StatusUnauthorized,
	1005,
	"用户身份过期",
)

var UserRegisterError = New(http.StatusBadRequest,
	1006,
	"用户注册失败",
)

var UserAlreadyExists = New(
	http.StatusBadRequest,
	utils.ToInt(i18n.I18N.T("Status.UserAlreadyExists")),
	i18n.I18N.T("Error.UserAlreadyExists"),
	// "用户已经存在",
)

var UserLoginFail = New(
	http.StatusBadRequest,
	1008,
	"用户登录失败",
)

var UserNotExistsErr = New(
	http.StatusBadRequest,
	1009,
	"用户不存在",
)

var UserPasswordNotCorrectErr = New(
	http.StatusBadRequest,
	1010,
	"用户密码错误",
)

var LimitPermissionError = New(
	http.StatusForbidden,
	1011,
	"权限受限,无法操作",
)

var InernalError = New(
	http.StatusInternalServerError,
	1012,
	"服务器内部错误",
)

var CommentError = New(
	http.StatusBadRequest,
	1013,
	"comment error",
)

var CommentNotFound = New(
	http.StatusNotFound,
	1014,
	"comment not found",
)

var CommentOperatorError = New(
	http.StatusBadRequest,
	1015,
	"评论操作失败",
)

var UnsupportedFileType = New(
	http.StatusBadRequest,
	1016,
	"Unsupported file type",
)
