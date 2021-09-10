package models

import (
	"gorm.io/gorm"
)

// Post represents user model
type Post struct {
	gorm.Model
	Title      string `gorm:"type:VARCHAR(255)" form:"title" json:"title" binding:"required,alphanum"`
	Slug       string `gorm:"type:VARCHAR(255);unique_index" form:"slug" json:"slug" binding:"alphanum,unique"`
	Content    string `form:"content" json:"content" binding:"required,alphanum"`
	UserID     uint
	User       User
	CategoryID uint
	Category   Category
	Tags       []*Tag `gorm:"many2many:post_tags;"`
}
