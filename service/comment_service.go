package service

import (
	"cabbage-server/dao"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"errors"

	"gorm.io/gorm"
)

func CreatePostComment(userid int64, comment *dto.CommentDTO) error {
	_, err := dao.CreateComment(userid, int64(comment.Post),comment.Content)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return internal.RequestParamsNotValidError
		}
		return internal.CommentError
	}
	return nil
}

func CommentView() {

}

func CommentOperator(userid int64, operator *dto.CommentOperatorDTO) error {
	_, err := dao.CommentOperator(userid, int64(operator.Cid), operator.Op)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return internal.CommentNotFound
		}
		return internal.CommentOperatorError
	}
	return nil
}
