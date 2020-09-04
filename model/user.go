package model

import "gorm.io/gorm"

// User ユーザ
type User struct {
	gorm.Model
	Name string
}
