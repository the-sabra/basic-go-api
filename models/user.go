package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Books    []Book `json:"books" gorm:"foreignKey:UserID"`
	Role     string `json:"role" gorm:"default:customer"`
}
