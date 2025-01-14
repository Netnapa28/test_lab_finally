package entity

import(
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username string `valid:"required~Username is required."`
	Password string
	Email    string `valid:"email~Email is invalid."`
}