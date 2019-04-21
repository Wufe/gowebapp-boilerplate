package models

import "github.com/jinzhu/gorm"

// User - The main user
type User struct {
	gorm.Model
	Email    string
	Password string

	Roles []*Role `gorm:"many2many:users_roles"`
}
