package models

import "github.com/jinzhu/gorm"

// Role - The role which gets associated to a user
type Role struct {
	gorm.Model
	Name        string
	Description string
	Users       []*User `gorm:"many2many:users_roles"`
}
