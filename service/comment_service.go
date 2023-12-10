package service

import (
	"cabbage-server/dao"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/model"
	"errors"

	"gorm.io/gorm"
)

func CreatePostComment(userid int64, comment *dto.CommentDTO) error {
	_, err := dao.CreateComment(userid, int64(comment.Post), comment.Content)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return internal.RequestParamsNotValidError
		}
		return internal.CommentError
	}
	return nil
}

func CommentView(postid int64) ([]*model.Comment, error) {
	commentList, err := dao.FindCommentByPost(postid)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.CommentNotFound
	}
	return commentList, nil
}

func ReplyPostComment(userid int64, postid int64, commentid int64, content string) error {
	comment, err := dao.CreateComment(userid, postid, content)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return internal.RequestParamsNotValidError
		}
		return internal.CommentError
	}
	comment.Parent = commentid
	err = dao.UpdateComment(comment)
	if err != nil {
		return internal.CommentError
	}
	return nil
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
