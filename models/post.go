package models

import "database/sql"

type Post struct {
	Id          int64        `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	PostId      uint64       `gorm:"column:post_id;type:bigint(20);comment:帖子id;NOT NULL" json:"post_id"`
	Title       string       `gorm:"column:title;type:varchar(128);comment:标题;NOT NULL" json:"title"`
	Content     string       `gorm:"column:content;type:varchar(8192);comment:内容;NOT NULL" json:"content"`
	AuthorId    uint64       `gorm:"column:author_id;type:bigint(20);comment:作者的用户id;NOT NULL" json:"author_id"`
	CommunityId uint64       `gorm:"column:community_id;type:bigint(20);comment:所属社区;NOT NULL" json:"community_id"`
	Status      int          `gorm:"column:status;type:tinyint(4);default:1;comment:帖子状态;NOT NULL" json:"status"`
	CreateTime  sql.NullTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime  sql.NullTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}

func (m *Post) TableName() string {
	return "post"
}

type ApiPostDetail struct {
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"`
}
