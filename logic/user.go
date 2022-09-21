package logic

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	var exist bool
	exist, err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("用户已存在")
	}
	// 生成UID
	userID := snowflake.GenID()

	// 构造一个User实例
	// 密码加密
	// 保存进数据库
	mysql.InsertUser()
}
