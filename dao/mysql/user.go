package mysql

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Id         int64        `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int64        `gorm:"column:user_id;type:bigint(20);NOT NULL" json:"user_id"`
	Username   string       `gorm:"column:username;type:varchar(64);NOT NULL" json:"username"`
	Password   string       `gorm:"column:password;type:varchar(64);NOT NULL" json:"password"`
	Email      string       `gorm:"column:email;type:varchar(64)" json:"email"`
	Gender     int          `gorm:"column:gender;type:tinyint(4);default:0;NOT NULL" json:"gender"`
	CreateTime sql.NullTime `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime sql.NullTime `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"update_time"`
}

func (m *User) TableName() string {
	return "user"
}

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (bool, error) {
	var result User
	find := db.Select("id").Find(&result, User{Username: username})

	// 查不到数据
	is := errors.Is(find.Error, gorm.ErrRecordNotFound)
	if is {
		return false, find.Error
	}
	return true, nil
}

func InsertUser() {
	// 执行SQL语句入库
}
