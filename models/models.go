package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `validate:"min=1,max=50"`
	Email    string `validate:"email" gorm:"unique"`
	Password string `validate:"min=8,max=100"`
}
