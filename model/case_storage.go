package model

import (
	"time"

	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type Case struct {
	gorm.Model
	CaseName string    `json:"caseName" gorm:"column:case_name" form:"caseName"`
	CaseDate time.Time `json:"-" gorm:"column:case_date"`
	// As a middleware for time conversion.
	CaseDateString string `json:"caseDate" form:"caseDate" gorm:"-"`

	// The logical foreign key.
	StaticAnalysisID  uint   `json:"staticAnalysisID" gorm:"column:static_analysis_id" form:"staticAnalysisID"`
	DynamicAnalysisID uint   `json:"dynamicAnalysisID" gorm:"column:dynamic_analysis_id" form:"dynamicAnalysisID"`
	ReportID          uint   `json:"reportID" gorm:"column:report_id" form:"reportID"`
	UploaderID        uint   `json:"uploaderID" gorm:"column:uploader_id" form:"uploaderID"`
	UploaderUserName  string `json:"uploaderUsername" gorm:"column:uploader_username" form:"uploaderUserName"`
}

// Hook function.
func (c *Case) BeforeCreate(db *gorm.DB) (err error) {
	if c.CaseDateString == "" {
		c.CaseDateString = util.GetTimestampNow()
	}
	c.CaseDate, err = util.FormatTimeString(c.CaseDateString)
	return
}
