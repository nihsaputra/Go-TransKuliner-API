package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	PhoneNumber string `gorm:"column:phone_number"`
}
