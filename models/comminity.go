package models

import "time"

type Community struct {
	Id            uint64    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	CommunityId   uint64    `gorm:"column:community_id;type:int(10) unsigned;NOT NULL" json:"community_id"`
	CommunityName string    `gorm:"column:community_name;type:varchar(128);NOT NULL" json:"community_name"`
	Introduction  string    `gorm:"column:introduction;type:varchar(256);NOT NULL" json:"introduction"`
	CreateTime    time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
}

func (m *Community) TableName() string {
	return "community"
}

type CommunityDetail struct {
	CommunityID   uint64    `json:"community_id" db:"community_id"`
	CommunityName string    `json:"community_name" db:"community_name"`
	Introduction  string    `json:"introduction,omitempty" db:"introduction"` // omitempty 当Introduction为空时不展示
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}
