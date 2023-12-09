package dao

import (
	"cabbage-server/db"
	"cabbage-server/dto"
	"cabbage-server/model"
)

const (
	LIKE = 1 << iota
	DISS
	FAVORITE
	SHARE
)

func CreatePost(author string, post *dto.PostDTO) (*model.Post, error) {
	_post := &model.Post{
		Author:   author,
		Title:    post.Title,
		Content:  post.Content,
		Like:     0,
		Diss:     0,
		Share:    0,
		Favorite: 0,
	}
	err := db.DB.Model(&model.Post{}).Create(_post).Error
	if err != nil {
		return nil, err
	} else {
		return _post, nil
	}
}

func FindPostByTitle(title string, page int, size int) {

}

func FindPostByAuthor(author string, page int, size int) {

}

func DeletePost() {

}

func FindPostById(pid int64) (*model.Post, error) {
	post := &model.Post{}
	err := db.DB.
		Model(&model.Post{}).
		Where("id = ?", pid).
		First(post).
		Error
	if err != nil {
		return nil, err
	} else {
		return post, nil
	}
}

func OperatorPost(postid int64, userid int64, opcode int) error {
	post, err := FindPostById(postid)
	if err != nil {
		return err
	}
	switch opcode {
	case LIKE:
		post.Like += 1
	case DISS:
		post.Diss += 1
	case SHARE:
		post.Share += 1
	case FAVORITE:
		post.Favorite += 1
	}
	_, err = CreatePostOperator(userid, postid, opcode)
	if err != nil {
		return err
	}
	err = db.DB.Model(&model.Post{}).Save(post).Error
	if err != nil {
		return err
	}
	return nil
}

func CreatePostOperator(uid int64, postid int64, opcode int) (*model.PostOperator, error) {
	CO := model.PostOperator{
		UserID: uid,
		PostID: postid,
		OPCode: opcode,
	}
	err := db.DB.Model(&model.CommentOperator{}).Save(&CO).Error
	if err != nil {
		return nil, err
	} else {
		return &CO, nil
	}
}