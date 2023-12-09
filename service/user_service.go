package service

import (
	"cabbage-server/dao"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateAccount 创建新用户服务
func CreateAccount(user *dto.SignupDTO) (*model.User, error) {
	_user := &model.User{
		UserId: uuid.New(),
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}
	err := dao.CreateAccount(_user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.UserNotFoundError
		} else {
			return nil, internal.UserRegisterError
		}
	}
	_profile := &model.UserProfile{
		UserId:   _user.ID,
		Twitter:  "",
		Mastodon: "",
		Facebook: "",
		Youtobe:  "",
		Gmail:    "",
		Github:   "",
		Insgram:  "",
		Telegram: "",
	}
	_, _err := dao.CreateProfile(_profile)
	if _err != nil {
		if !errors.Is(_err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.UserRegisterError
	}
	return _user, nil
}

// GetUserProfile 获取用户信息服务
func GetUserProfile(email string) (*model.User, error) {
	user, err := dao.FindUserByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.RecordNotFoundError
		} else {
			return nil, internal.InernalError
		}
	}
	return user, nil
}

func FindUserByName(name string) (*model.User, error) {
	user, err := dao.FindUserByName(name)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.UserNotFoundError
	}
	return user, nil
}

func CheckNickName(name string) error {
	_, err := dao.FindUserByName(name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return internal.UserAlreadyExists
	}
	return nil
}

func ProfileShare(user string) (*model.UserProfile, error) {
	_user, err := dao.FindUserByName(user)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.RecordNotFoundError
	}
	profile, err := dao.FindProfileByUID(int64(_user.ID))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.RecordNotFoundError
	}
	return profile, nil
}
