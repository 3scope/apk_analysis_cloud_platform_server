package model

import "gorm.io/gorm"

type ApkStorage struct {
	gorm.Model
	AppName    string `json:"appName" gorm:"column:app_name"`
	UploadedBy string `json:"uploadedBy" gorm:"column:uploaded_by"`
	// To store the path where the apk installer is.
	ApkPath string `json:"-" gorm:"column:apk_path"`
}
