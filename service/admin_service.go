package service

import (
	"cabbage-server/dao"
	"cabbage-server/internal"
	"cabbage-server/model"
	"errors"

	"gorm.io/gorm"
)

func AddPlatform(names ...string) error {
	for _, name := range names {
		_, err := dao.FindPlatformByName(name)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return internal.RecordUpdateError
			}
			_, err = dao.AddNewPlatform(name)
			if err != nil {
				return internal.RecordUpdateError
			}
		}
	}
	return nil
}

func GetAllPlatform() ([]string, error) {
	list, err := dao.GetAllPlatform()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.RecordNotFoundError
	}
	return list, nil
}

func Count(month int) (map[string]any, error) {
	commentMonth, err := dao.CountNewCommentOfMonth(month)
	if err != nil {
		return nil, err
	}
	userMonth, err := dao.CountNewUserOfMonth(month)
	if err != nil {
		return nil, err
	}
	postMonth, err := dao.CountNewPostOfMonth(month)

	if err != nil {
		return nil, internal.InernalError
	}

	commentToday, err := dao.CountNewCommentOfToday()
	if err != nil {
		return nil, internal.InernalError
	}
	UserToday, err := dao.CountNewUserOfToday()
	if err != nil {
		return nil, err
	}

	postToday, err := dao.CountNewPostOfToday()
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"month": map[string]any{
			"user":    userMonth,
			"post":    postMonth,
			"comment": commentMonth,
		},
		"today": map[string]any{
			"user":    UserToday,
			"post":    postToday,
			"comment": commentToday,
		},
	}, nil

}

func CountCommentOfToday() (int64, error) {
	count, err := dao.CountNewCommentOfToday()
	if err != nil {
		return -1, internal.InernalError
	}
	return count, nil
}

func CountPostOfToday() (int64, error) {
	count, err := dao.CountNewPostOfToday()
	if err != nil {
		return -1, internal.InernalError
	}
	return count, nil
}

func CountUserOfToday() (int64, error) {
	count, err := dao.CountNewUserOfToday()
	if err != nil {
		return -1, internal.InernalError
	}
	return count, nil
}

func CountCommentOfMonth(month int) ([]*model.Counts, error) {
	result, err := dao.CountNewCommentOfMonth(month)
	if err != nil {
		return nil, internal.InernalError
	}
	return result, nil
}

func CountPostOfMonth(month int) ([]*model.Counts, error) {
	result, err := dao.CountNewPostOfMonth(month)
	if err != nil {
		return nil, internal.InernalError
	}
	return result, nil
}

func CountUserOfMonth(month int) ([]*model.Counts, error) {
	result, err := dao.CountNewPostOfMonth(month)
	if err != nil {
		return nil, internal.InernalError
	}
	return result, nil
}
