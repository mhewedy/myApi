package profile

import (
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Name  string
	Email string
}
