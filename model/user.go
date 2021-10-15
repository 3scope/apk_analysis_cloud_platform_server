package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"column:username"`
	RealName    string `json:"realName" gorm:"column:real_name"`
	Role        string `json:"role" gorm:"column:role"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number"`
	Description string `json:"description" gorm:"column:description"`
	// To keep password string private.
	Password string `json:"-,omitempty" gorm:"column:password"`
}
