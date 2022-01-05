package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"column:username" form:"username"`
	RealName    string `json:"realName" gorm:"column:real_name" form:"realName"`
	Role        string `json:"role" gorm:"column:role" form:"role"`
	Email       string `json:"email" gorm:"column:email" form:"email"`
	Description string `json:"description" gorm:"column:description" form:"description"`

	// To keep password string private.
	Password string `json:"-" gorm:"column:password" form:"password"`
}
