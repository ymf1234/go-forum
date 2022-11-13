package logic

import (
	"go.uber.org/zap"
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
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("GetPostList mysql.GetUserByID failed",
				zap.Int64("postID", post.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		mysql.GetCommunityById(post.CommunityId)
	}
	return
}

func GetPostListNew() {

}
