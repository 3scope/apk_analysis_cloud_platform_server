package model

import "gorm.io/gorm"

type DynamicAnalysis struct {
	gorm.Model
	AnalysisData string `json:"analysisData" gorm:"column:analysis_data" form:"analysisData"`

	// The logical foreign key.
	VideoID uint `json:"videoID" gorm:"column:video_id" form:"videoID"`
	CaseID  uint `json:"caseID" gorm:"column:case_id" form:"caseID"`
}
