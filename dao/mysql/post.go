package mysql

import (
	"errors"
	"gorm.io/gorm"
	"web_app/models"
)

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	filed := []string{
		"post_id",
		"title",
		"content",
		"author_id",
		"community_id",
		"create_time",
	}
	find := db.Select(filed).Limit(int(size)).Offset(int((page - 1) * size)).Find(&posts)

	if find.Error != nil {
		// 查不到数据
		is := errors.Is(find.Error, gorm.ErrRecordNotFound)
		if is {
			return nil, errors.New("未查询出数据")
		}
		return nil, find.Error
	}
	return posts, nil
}
