package routes

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"
)

func Setup() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode) // 发布模式
	r := gin.New()

	// 日志中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/sign-up", controller.SignUpHandler)
	// 登录
	r.POST("/login", controller.LoginHandler)
	// 刷新token
	r.GET("/refresh_token", controller.RefreshTokenHandle)

	// 列表显示
	r.GET("/posts", controller.PostListHandler)   // 分页展示帖子列表
	r.GET("/posts2", controller.PostList2Handler) // 根据时间或者分数排序分页展示帖子列表

	// 分类
	r.GET("/community", controller.CommunityHandler)           //获取分类社区列表
	r.GET("/community/:id", controller.CommunityDetailHandler) //根据ID查找社区详情

	// 帖子
	r.GET("/post/:id", controller.PostDetailHandler) // 查询帖子详情

	r.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
	{
		r.POST("/post", controller.CreatePostHandler) // 创建帖子

		r.POST("/vote", controller.VoteHandler) // 投票

		r.POST("/comment", controller.CommentHandler)
		r.GET("/comment", controller.CommentListHandler)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	pprof.Register(r) // 注册pprof相关路由
	// 找不到页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
