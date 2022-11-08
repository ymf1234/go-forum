package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"web_app/models"
)

const secret = "explan.com"

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (bool, error) {
	var result models.User
	find := db.Select("id").Find(&result, models.User{Username: username})

	// 查不到数据
	is := errors.Is(find.Error, gorm.ErrRecordNotFound)
	if is {
		return false, find.Error
	}
	return true, nil
}

// InsertUser 新增用户
func InsertUser(user *models.User) (err error) {
	// 对密码加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	result := db.Create(&user)
	// user.id             // 返回插入数据的主键
	err = result.Error // 返回 error
	//result.RowsAffected // 返回插入记录的条数
	return
}

// encryptPassword md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(p *models.User) (err error) {
	originPassword := p.Password // 记录一下原始密码(用户登录的密码)
	find := db.Find(&p, models.User{Username: p.Username})

	// 查不到数据
	is := errors.Is(find.Error, gorm.ErrRecordNotFound)
	if is {
		return find.Error
	}

	if originPassword != encryptPassword(p.Password) {
		return errors.New("密码错误")
	}
	return nil
}
