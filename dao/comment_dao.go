package dao

import (
	"cabbage-server/db"
	"cabbage-server/dto"
	"cabbage-server/model"
)

func CreateComment(userid int64, comment *dto.CommentDTO) (*model.Comment, error) {
	_comment := &model.Comment{
		Content: comment.Content,
		UserID:  userid,
		PostId:  int64(comment.Post),
	}
	err := db.DB.Model(&model.Comment{}).Create(_comment).Error
	if err != nil {
		return nil, err
	} else {
		return _comment, nil
	}
}

func DeleteComment(commentId int64) error {
	comment, err := FindCommentById(commentId)
	if err != nil {
		return err
	} else {
		err = db.DB.Model(&model.Comment{}).Delete(comment).Error
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

func CreateReplyComment(userid int64, commentid int64, content string) (*model.Comment, error) {
	comment := &model.Comment{
		UserID:  userid,
		Parent:  commentid,
		Content: content,
	}
	err := db.DB.Model(&model.Comment{}).Create(comment).Error
	if err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}

func FindCommentById(commentid int64) (*model.Comment, error) {
	comment := &model.Comment{}
	err := db.DB.Model(&model.Comment{}).Where("id = ?").First(comment).Error
	if err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}

func FindCommentReply(parent int64, page int, size int) ([]*model.Comment, error) {
	var replyList []*model.Comment
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	} else if size >= 20 {
		size = 20
	}
	err := db.DB.
		Model(&model.Comment{}).
		Where("parent = ?", parent).
		Limit(size).
		Offset((page - 1) * size).
		Find(&replyList).
		Error
	if err != nil {
		return nil, err
	} else {
		return replyList, nil
	}

}

func FindCommentByPost(postid int64) ([]*model.Comment, error) {
	var commentList []*model.Comment
	err := db.DB.Model(&model.Comment{}).Where("post_id = ?", postid).Find(commentList).Error
	if err != nil {
		return nil, err
	} else {
		return commentList, nil
	}
}
