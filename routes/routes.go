package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/logger"
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

	r.GET("/posts", controller.PostListHandler)   // 分页展示帖子列表
	r.GET("/posts2", controller.PostList2Handler) // 根据时间或者分数排序分页展示帖子列表

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
