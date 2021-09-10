package models

import (
	"gorm.io/gorm"
)

// Post represents user model
type Tag struct {
	gorm.Model
	Name  string  `gorm:"type:VARCHAR(100)" form:"name" json:"name" binding:"required,alphanum"`
	Slug  string  `gorm:"type:VARCHAR(100);unique_index" form:"slug" json:"slug" binding:"alphanum,unique"`
	Posts []*Post `gorm:"many2many:post_tags;"`
}
