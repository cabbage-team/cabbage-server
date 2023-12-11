package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// CreateAccount 创建用户数据库操作
func CreateAccount(username, password string, email string) error {
	user := &model.User{
		UserId:   uuid.New(),
		Name:     username,
		Password: password,
		Email:    email,
	}
	err := db.DB.Model(&model.User{}).Create(user).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

// GetUserProfile 获取用户信息数据库操作
func FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Model(&model.User{}).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Model(&model.User{}).
		Where("name = ?", name).
		First(user).
		Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func CountNewUserOfMonth(month int) ([]*model.Counts, error) {
	timedate := time.Now()
	var results []*model.Counts
	year, _, _ := timedate.Date()
	dateString := strings.Join([]string{fmt.Sprintf("%d", year), fmt.Sprintf("%d", month), "01"}, "-")
	nextMonth := strings.Join([]string{fmt.Sprintf("%d", year), fmt.Sprintf("%d", month+1), "01"}, "-")
	err := db.DB.Model(&model.User{}).
		Select("DATE(created_at) as `date`", "count(*) as counts").
		Where("created_at >= ? AND created_at < ?", dateString, nextMonth).
		Order("DATE(created_at)").
		Group("date(created_at)").
		Find(&results).Error
	if err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func CountNewUserOfToday() (int64, error) {
	var count int64
	err := db.DB.Model(&model.User{}).
		Where("DATE(created_at) = CURDATE()").
		Count(&count).
		Error
	if err != nil {
		return -1, err
	}
	return count, nil
}
