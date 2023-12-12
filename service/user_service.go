package service

import (
	"cabbage-server/dao"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/model"
	"errors"

	"gorm.io/gorm"
)

// CreateAccount 创建新用户服务
func CreateAccount(user *dto.SignupDTO) (*model.User, error) {
	_user, err := dao.CreateAccount(user.Name, user.Password, user.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.UserNotFoundError
		} else {
			return nil, internal.UserRegisterError
		}
	}

	list, err := dao.GetAllPlatform()
	if err != nil {
		return nil, internal.UserRegisterError
	}
	for _, platform := range list {
		_, _err := dao.CreateProfile(&model.UserProfile{
			UserId:   _user.ID,
			Platform: platform,
			URI:      "",
		})
		if _err != nil {
			if !errors.Is(_err, gorm.ErrRecordNotFound) {
				return nil, internal.InernalError
			}
			return nil, internal.UserRegisterError
		}
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

func ProfileShare(user string) (map[string]string, error) {
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
	result := map[string]string{}
	for _, val := range profile {
		result[val.Platform] = val.URI
	}

	return result, nil
}

// 用户黑名单
func UserBlackList(userid int64, page, size int) ([]*model.User, error) {
	list, err := dao.GetUserBlackList(userid, page, size)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.UserNotFoundError
	}
	return list, nil
}
// 用户关注列表
func UserFollowList(userid int64, page, size int) ([]*model.User, error) {
	list, err := dao.UserFollows(userid, page, size)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.InernalError
		}
		return nil, internal.UserNotFoundError
	}
	return list, nil
}
