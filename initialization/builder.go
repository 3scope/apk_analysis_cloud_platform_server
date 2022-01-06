package initialization

import (
	"github.com/sanscope/apk_analysis_cloud_platform_server/handler"
	"gorm.io/gorm"
)

type InitializationList struct {
	DB                     *gorm.DB
	UserHandler            *handler.UserHandler
	CaseHandler            *handler.CaseHandler
	StaticAnalysisHandler  *handler.StaticAnalysisHandler
	DynamicAnalysisHandler *handler.DynamicAnalysisHandler
	VideoHandler           *handler.VideoHandler
	ReportHandler          *handler.ReportHandler
}

func InitializationBuilder() *InitializationList {
	return new(InitializationList)
}
