package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

// GetCommunityList 查询分类社区列表
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetailByID
func GetCommunityDetailByID(id uint64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityByID(id)
}
