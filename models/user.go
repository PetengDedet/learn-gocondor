// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"gorm.io/gorm"
)

// User represents user model
type User struct {
	gorm.Model
	Username string `gorm:"type:VARCHAR(255);unique_index;" json:"username"`
	Name     string `gorm:"type:VARCHAR(255);"`
	Email    string `gorm:"type:VARCHAR(255);unique_index;"`
	Password string `gorm:"type:VARCHAR(255);"`
	Posts    []Post
}
