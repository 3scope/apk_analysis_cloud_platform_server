package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	VideoName string `json:"videoName" gorm:"column:video_name" form:"videoName"`
	VideoTime string `json:"videoTime" gorm:"column:video_time" form:"videoTime"`
}
