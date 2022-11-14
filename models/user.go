package models

import "database/sql"

type User struct {
	Id           int64        `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	UserId       uint64       `gorm:"column:user_id;type:bigint(20);NOT NULL" json:"user_id"`
	Username     string       `gorm:"column:username;type:varchar(64);NOT NULL" json:"username"`
	Password     string       `gorm:"column:password;type:varchar(64);NOT NULL" json:"password"`
	Email        string       `gorm:"column:email;type:varchar(64)" json:"email"`
	Gender       int          `gorm:"column:gender;type:tinyint(4);default:0;NOT NULL" json:"gender"`
	CreateTime   sql.NullTime `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime   sql.NullTime `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"update_time"`
	AccessToken  string
	RefreshToken string
}

func (m *User) TableName() string {
	return "user"
}
