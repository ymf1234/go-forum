package logic

import (
	"errors"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	var exist bool
	exist, err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("用户已存在")
	}
	// 生成UID
	userID, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("snowflake.GenID() failed", zap.Error(err))
		return
	}
	// 构造一个User实例
	user := &models.User{
		UserId:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}

// Login 登陆
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}

	aToken, rToken, err := jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return user, err
	}
	user.AccessToken = aToken
	user.RefreshToken = rToken
	return user, nil
}
