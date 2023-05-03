package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Term       string
	Definition string
}

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
