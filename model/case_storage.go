package model

import (
	"time"

	"gorm.io/gorm"
)

type Case struct {
	gorm.Model
	CaseName string    `json:"caseName" gorm:"column:case_name" form:"caseName"`
	CaseDate time.Time `json:"caseDate" gorm:"column:case_date" form:"caseDate"`

	// The logical foreign key.
	StaticAnalysisID  uint   `json:"staticAnalysisID" gorm:"column:static_analysis_id" form:"staticAnalysisID"`
	DynamicAnalysisID uint   `json:"dynamicAnalysisID" gorm:"column:dynamic_analysis_id" form:"dynamicAnalysisID"`
	ReportID          uint   `json:"reportID" gorm:"column:report_id" form:"reportID"`
	UploaderID        uint   `json:"uploaderID" gorm:"column:uploader_id" form:"uploaderID"`
	UploaderUserName  string `json:"uploaderUsername" gorm:"column:uploader_username" form:"uploaderUserName"`
}
