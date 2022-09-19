package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/models"
)

func SignUpHandler(c *gin.Context) {
	// 1、参数校验
	var p models.ParamSignUp
	err := c.ShouldBindJSON(&p)
	if err != nil {
		// 返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	fmt.Println(p)
	// 2、业务处理
	logic.SignUp()
	// 3、返回响应
	c.JSON(http.StatusOK, "ok")
}
