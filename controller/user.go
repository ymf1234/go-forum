package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		// 判断err是不是 validator.ValidationErrors 类型
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errors.Translate(trans)), //翻译错误
		})
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	/*if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}*/
	// 2、业务处理
	logic.SignUp(p)
	// 3、返回响应
	c.JSON(http.StatusOK, "ok")
}
