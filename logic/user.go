package logic

import (
	"web_app/dao/mysql"
	"web_app/pkg/snowflake"
)

func SignUp() {
	// 判断用户存不存在
	mysql.QueryUserByUsername()
	// 生成UID
	snowflake.GenID()
	// 密码加密
	// 保存进数据库
	mysql.InsertUser()
}
