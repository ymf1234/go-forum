package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		// 查询作者信息
		mysql.GetUserByID(post.AuthorId)
	}
	return
}

func GetPostListNew() {

}
