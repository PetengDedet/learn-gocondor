package models

import (
	"time"

	"gorm.io/gorm"
)

// Post represents user model
type Category struct {
	// gorm.Model
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time     `gorm:"index" json:"-"`
	UpdatedAt *time.Time     `gorm:"index" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"type:VARCHAR(100)" form:"name" json:"name" binding:"required,alphanum"`
	Slug      string         `gorm:"type:VARCHAR(100);unique_index" form:"slug" json:"slug" binding:"alphanum,unique"`
	ParentID  *uint          `json:"-"`
	Parent    *Category      `json:"parent"`
	Posts     []Post         `json:"posts"`
}
