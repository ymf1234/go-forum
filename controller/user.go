package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strings"
	"web_app/logic"
	"web_app/models"
	"web_app/pkg/jwt"
)

// SignUpHandler 注册请求
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
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errors.Translate(trans)))
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
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp  failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "注册失败")
		return
	}
	// 3、返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录请求
func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))

		return
	}
	// 业务逻辑处理
	user, err := logic.Login(p)

	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "用户名或密码错误")
		return
	}
	// 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserId),
		"user_name":     user.Username,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

// RefreshTokenHandle 刷新accessToken
func RefreshTokenHandle(c *gin.Context) {
	rt := c.Query("refresh_token")

	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}

	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	atoken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		zap.L().Error("logic.RefreshToken failed", zap.Error(err))
	}
	ResponseSuccess(c, gin.H{
		"access_token":  atoken,
		"refresh_token": rToken,
	})
}
