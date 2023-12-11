package dao

import (
	"cabbage-server/db"
	"cabbage-server/model"
	"fmt"
	"strings"
	"time"
)

// 分页读取话题标签
func ReadTags(page, size int) (*[]model.Tag, error) {
	var tags []model.Tag
	err := db.DB.
		Model(&model.Tag{}).
		Limit(size).
		Offset((page - 1) * size).
		Find(&tags).
		Error
	if err != nil {
		return nil, err
	} else {
		return &tags, nil
	}
}

// 查找指定 id的话题
func FindTagByIds(tags ...int64) ([]*model.Tag, error) {
	_tags := []*model.Tag{}
	err := db.DB.Model(&model.Tag{}).
		Omit("created_at", "deleted_at", "updated_at").
		Where("id in (?)", tags).
		Find(&_tags).
		Error
	if err != nil {
		return nil, err
	}
	return _tags, nil
}

// 统计某月新增话题数
func CountNewTagOfMonth(month int) ([]*model.Counts, error) {
	timedate := time.Now()
	var results []*model.Counts
	year, _, _ := timedate.Date()
	dateString := strings.Join([]string{fmt.Sprintf("%d", year), fmt.Sprintf("%d", month), "01"}, "-")
	nextMonth := strings.Join([]string{fmt.Sprintf("%d", year), fmt.Sprintf("%d", month+1), "01"}, "-")
	err := db.DB.Model(&model.Tag{}).
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
// 统计当天新增话题数
func CountNewTagOfToday() (int64, error) {
	var count int64
	err := db.DB.Model(&model.Tag{}).
		Where("DATE(created_at) = CURDATE()").
		Count(&count).
		Error
	if err != nil {
		return -1, err
	}
	return count, nil
}
