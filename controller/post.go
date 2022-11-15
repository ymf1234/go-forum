package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 1、获取参数及校验
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		zap.L().Debug("c.ShouldBindJSON(post)", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	// 获取作者ID， 当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorId = userID

	// 创建帖子
	logic.CreatePost(&post)
}

// PostListHandler 分页获取帖子列表
func PostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("PostListHandler err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

func PostList2Handler(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	//c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("PostList2Handler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取数据
	logic.GetPostListNew(p)
}

// PostDetailHandler
func PostDetailHandler(c *gin.Context) {
	// 1、获取参数(从URL中获取帖子的id)
	postIdStr := c.Param("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
	}

	// 根据id取出id帖子数据(查数据库)
	post, err := logic.GetPostById(postId)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	// 3、返回响应
	ResponseSuccess(c, post)
}
