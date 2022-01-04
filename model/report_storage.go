package model

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ReportName string `json:"reportName" gorm:"report_name" form:"reportName"`

	// The logical foreign key
	CaseID uint `json:"caseID" gorm:"case_id" form:"caseID"`
}
