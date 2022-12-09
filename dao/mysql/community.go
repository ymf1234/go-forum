package mysql

import (
	"errors"
	"gorm.io/gorm"
	"web_app/models"
)

// GetCommunityList 获取全部数据
func GetCommunityList() (communityList []*models.Community, err error) {
	field := []string{
		"community_id",
		"community_name",
	}
	find := db.Select(field).Find(&communityList)
	if find.Error != nil {
		// 查不到数据
		is := errors.Is(find.Error, gorm.ErrRecordNotFound)
		if is {
			return nil, errors.New("未查询出数据")
		}
		return nil, find.Error
	}
	return communityList, nil
}

// GetCommunityNameByID 获取指定数据
func GetCommunityNameByID(communityId uint64) (community *models.Community, err error) {
	field := []string{
		"community_id",
		"community_name",
	}
	first := db.Select(field).Where("community_id = ?", communityId).First(&community)

	if first.Error != nil {
		// 查不到数据
		is := errors.Is(first.Error, gorm.ErrRecordNotFound)
		if is {
			return nil, errors.New("未查询出数据")
		}
		return nil, first.Error
	}
	return community, nil
}

// GetCommunityByID 根据ID查询分类社区详情
func GetCommunityByID(id uint64) (community *models.CommunityDetail, err error) {
	//community = new(models.CommunityDetail)
	field := []string{
		"community_id",
		"community_name",
		"introduction",
		"create_time",
	}
	find := db.Select(field).Where("community_id = ?", id).Find(&community)

	if find.Error != nil {
		// 查不到数据
		is := errors.Is(find.Error, gorm.ErrRecordNotFound)
		if is {
			return nil, errors.New("未查询出数据")
		}
		return nil, find.Error
	}
	return community, nil
}
