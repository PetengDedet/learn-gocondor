// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents user model
type User struct {
	// gorm.Model
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time     `gorm:"index" json:"-"`
	UpdatedAt *time.Time     `gorm:"index" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"type:VARCHAR(255);unique_index;" json:"username"`
	Name      string         `gorm:"type:VARCHAR(255);" json:"name"`
	Email     string         `gorm:"type:VARCHAR(255);unique_index;" json:"email"`
	Password  string         `gorm:"type:VARCHAR(255);" json:"-"`
	Posts     []Post         `json:"posts"`
}
