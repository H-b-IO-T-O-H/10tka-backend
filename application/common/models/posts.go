package models

import "github.com/go-openapi/strfmt"

type Post struct {
	PostId   string          `gorm:"column:id" json:"post_id"`
	AuthorId string          `gorm:"column:author_id" json:"author_id" binding:"required"`
	TagType  string          `gorm:"column:tag_type" json:"tag_type" binding:"required"`
	Title    string          `gorm:"column:title" json:"title" binding:"required"`
	Content  string          `gorm:"column:content" json:"content" binding:"required"`
	IsEdited bool            `gorm:"column:is_edited" json:"is_edited"`
	Created  strfmt.DateTime `gorm:"column:created" json:"created"`
	Comments bool            `gorm:"column:comments" json:"comments"`
}
