package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
	Role     string
}
