package models

import "database/sql"

type Comment struct {
	Id         int64        `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	CommentId  uint64       `gorm:"column:comment_id;type:bigint(20) unsigned;NOT NULL" json:"comment_id"`
	Content    string       `gorm:"column:content;type:text;NOT NULL" json:"content"`
	PostId     int64        `gorm:"column:post_id;type:bigint(20);NOT NULL" json:"post_id"`
	AuthorId   int64        `gorm:"column:author_id;type:bigint(20);NOT NULL" json:"author_id"`
	ParentId   int64        `gorm:"column:parent_id;type:bigint(20);default:0;NOT NULL" json:"parent_id"`
	Status     uint         `gorm:"column:status;type:tinyint(3) unsigned;default:1;NOT NULL" json:"status"`
	CreateTime sql.NullTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime sql.NullTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"update_time"`
}

func (m *Comment) TableName() string {
	return "comment"
}
