package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
	"fmt"
	"strings"
	"time"
)

func CreateComment(userid int64, postid int64, content string) (*model.Comment, error) {
	_comment := &model.Comment{
		Content:  content,
		UserID:   userid,
		PostID:   postid,
		Like:     0,
		Diss:     0,
		Favorite: 0,
		Share:    0,
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
	err := db.DB.Model(&model.Comment{}).
		Omit("created_at", "deleted_at", "updated_at").
		Where("id = ?").
		First(comment).
		Error
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
		Omit("created_at", "deleted_at", "updated_at").
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

func UpdateComment(comment *model.Comment) error {
	err := db.DB.Model(&model.Comment{}).
		Save(comment).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func FindCommentByPost(postid int64) ([]*model.Comment, error) {
	var commentList []*model.Comment
	err := db.DB.Model(&model.Comment{}).
		Omit("created_at", "deleted_at", "updated_at").
		Where("post_id = ?", postid).
		Find(commentList).
		Error
	if err != nil {
		return nil, err
	} else {
		return commentList, nil
	}
}

func CreateCommentOperator(userid int64, cid int64, opcode int) (*model.CommentOperator, error) {
	cmtOP := &model.CommentOperator{
		CommentID: cid,
		OPCode:    opcode,
	}
	err := db.DB.Model(&model.CommentOperator{}).
		Create(cmtOP).Error
	if err != nil {
		return nil, err
	}
	return cmtOP, nil
}

func CommentOperator(userid int64, cid int64, opcode int) (*model.CommentOperator, error) {
	cmt, err := FindCommentById(cid)
	if err != nil {
		return nil, err
	}
	switch opcode {
	case LIKE:
		cmt.Like += 1
	case DISS:
		cmt.Diss += 1
	case SHARE:
		cmt.Share += 1
	case FAVORITE:
		cmt.Favorite += 1
	}
	err = UpdateComment(cmt)
	if err != nil {
		return nil, err
	}
	cmtOP, err := CreateCommentOperator(userid, cid, opcode)
	if err != nil {
		return nil, err
	}
	return cmtOP, nil
}

func CountNewCommentOfMonth(month int) ([]*model.Counts, error) {
	timedate := time.Now()
	var results []*model.Counts
	year, _, _ := timedate.Date()
	dateString := strings.Join([]string{fmt.Sprintf("%d", year), fmt.Sprintf("%d", month), "01"}, "-")
	nextMonth := strings.Join([]string{fmt.Sprintf("%d", year), fmt.Sprintf("%d", month+1), "01"}, "-")
	err := db.DB.Model(&model.Comment{}).
		Select("DATE(created_at) as `date`", "count(*) as counts").
		Where("created_at >= ? AND created_at < ?", dateString, nextMonth).
		Order("DATE(created_at)").
		Group("DATE(created_at)").
		Find(&results).Error
	if err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func CountNewCommentOfToday() (int64, error) {
	var count int64
	err := db.DB.Model(&model.Comment{}).
		Where("DATE(created_at) = CURDATE()").
		Count(&count).
		Error
	if err != nil {
		return -1, err
	}
	return count, nil
}
