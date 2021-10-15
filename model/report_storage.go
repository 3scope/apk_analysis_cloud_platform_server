package model

import "gorm.io/gorm"

type ReportStorage struct {
	gorm.Model
	ReportName string `json:"reportName" gorm:"report_name"`
	UploadedBy string `json:"uploadedBy" gorm:"column:uploaded_by"`
	AppName    string `json:"appName" gorm:"column:app_name"`
	// To store the path where the report is.
	ReportPath string `json:"-" gorm:"column:report_path"`
}
