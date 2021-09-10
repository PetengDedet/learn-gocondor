package models

import (
	"gorm.io/gorm"
)

// Post represents user model
type Category struct {
	gorm.Model
	Name     string `gorm:"type:VARCHAR(100)" form:"name" json:"name" binding:"required,alphanum"`
	Slug     string `gorm:"type:VARCHAR(100);unique_index" form:"slug" json:"slug" binding:"alphanum,unique"`
	ParentID *uint
	Parent   *Category
	Posts    []Post
}
