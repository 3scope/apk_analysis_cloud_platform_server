package model

import "gorm.io/gorm"

type VideoStorage struct {
	gorm.Model
	VideoName  string `json:"videoName" gorm:"column:video_name"`
	UploadedBy string `json:"uploadedBy" gorm:"column:uploaded_by"`
	AppName    string `json:"appName" gorm:"column:app_name"`
	// To store the path where the video is.
	VideoPath string `json:"-" gorm:"column:video_path"`
}
