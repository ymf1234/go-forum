package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	// 生成post_id(生成帖子ID)
	postID, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("snowflake.GenID() failed", zap.Error(err))
		return
	}

	post.PostId = postID
	// 创建帖子 保存到数据库
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}

	community, err := mysql.GetCommunityNameByID(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}

	// redis 存储帖子信息
	if err := redis.CreatePost(
		post.PostId,
		post.AuthorId,
		post.Title,
		TruncateByWords(post.Content, 120),
		community.CommunityId); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return
}

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
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityByID(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityId),
				zap.Error(err))
			continue
		}
		// 接口数据拼接
		postDetail := &models.ApiPostDetail{
			Post:            post,
			CommunityDetail: community,
			AuthorName:      user.Username,
		}

		data = append(data, postDetail)
	}
	return
}

func GetPostById(postID int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	// 查询帖子信息
	post, err := mysql.GetPostByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed",
			zap.Int64("postID", postID),
			zap.Error(err))
		return nil, err
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserByID(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed",
			zap.Uint64("postID", post.AuthorId),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityByID(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed",
			zap.Uint64("community_id", post.CommunityId),
			zap.Error(err))
		return
	}
	// 接口数据拼接
	data = &models.ApiPostDetail{
		Post:            post,
		CommunityDetail: community,
		AuthorName:      user.Username,
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	return
}

func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 根据请求参数的不同，执行不同的业务逻辑
	if p.CommunityID == 0 {
		// 查询所有
		GetPostList2(p)
	}

	return
}
