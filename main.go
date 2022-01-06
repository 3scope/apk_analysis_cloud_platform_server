package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/initialization"
	"github.com/spf13/viper"
)

func main() {
	// Gin default settings.
	router := gin.Default()
	router.Use(initialization.CORSSettings())
	gin.SetMode(viper.GetString("mode"))

	// Initial setup.
	instance := initialization.InitializationBuilder().InitializeDatabase().InitializeHandler()

	// User service router path.
	user := router.Group("/api-v1/user")
	{
		user.POST("/add", instance.UserHandler.AddHandler)
		user.GET("/get-total", instance.UserHandler.GetTotalHandler)
		user.GET("/list", instance.UserHandler.ListHandler)
		user.GET("/get-one", instance.UserHandler.GetOneHandler)
		user.GET("/is-exist", instance.UserHandler.IsExistHandler)
		user.DELETE("/delete", instance.UserHandler.DeleteHandler)
		user.PUT("/update", instance.UserHandler.UpdateHandler)
	}
	// Case service router path.
	caseInstance := router.Group("/api-v1/case")
	{
		caseInstance.POST("/add", instance.CaseHandler.AddHandler)
		caseInstance.GET("/get-total", instance.CaseHandler.GetTotalHandler)
		caseInstance.GET("/list", instance.CaseHandler.ListHandler)
		caseInstance.GET("/get-one", instance.CaseHandler.GetOneHandler)
		caseInstance.GET("/is-exist", instance.CaseHandler.IsExistHandler)
		caseInstance.DELETE("/delete", instance.CaseHandler.DeleteHandler)
		caseInstance.PUT("/update", instance.CaseHandler.UpdateHandler)
	}
	// Static analysis service router path.
	staticAnalysis := router.Group("/api-v1/static-analysis")
	{
		staticAnalysis.POST("/add", instance.StaticAnalysisHandler.AddHandler)
		staticAnalysis.GET("/get-total", instance.StaticAnalysisHandler.GetTotalHandler)
		staticAnalysis.GET("/list", instance.StaticAnalysisHandler.ListHandler)
		staticAnalysis.GET("/get-one", instance.StaticAnalysisHandler.GetOneHandler)
		staticAnalysis.GET("/is-exist", instance.StaticAnalysisHandler.IsExistHandler)
		staticAnalysis.DELETE("/delete", instance.StaticAnalysisHandler.DeleteHandler)
		staticAnalysis.PUT("/update", instance.StaticAnalysisHandler.UpdateHandler)
	}
	// Dynamic analysis service router path.
	dynamicAnalysis := router.Group("/api-v1/dynamicAnalysis")
	{
		dynamicAnalysis.POST("/add", instance.DynamicAnalysisHandler.AddHandler)
		dynamicAnalysis.GET("/get-total", instance.DynamicAnalysisHandler.GetTotalHandler)
		dynamicAnalysis.GET("/list", instance.DynamicAnalysisHandler.ListHandler)
		dynamicAnalysis.GET("/get-one", instance.DynamicAnalysisHandler.GetOneHandler)
		dynamicAnalysis.GET("/is-exist", instance.DynamicAnalysisHandler.IsExistHandler)
		dynamicAnalysis.DELETE("/delete", instance.DynamicAnalysisHandler.DeleteHandler)
		dynamicAnalysis.PUT("/update", instance.DynamicAnalysisHandler.UpdateHandler)
	}
	// Video service router path.
	video := router.Group("/api-v1/video")
	{
		video.POST("/add", instance.VideoHandler.AddHandler)
		video.GET("/get-total", instance.VideoHandler.GetTotalHandler)
		video.GET("/list", instance.VideoHandler.ListHandler)
		video.GET("/get-one", instance.VideoHandler.GetOneHandler)
		video.GET("/is-exist", instance.VideoHandler.IsExistHandler)
		video.DELETE("/delete", instance.VideoHandler.DeleteHandler)
		video.PUT("/update", instance.VideoHandler.UpdateHandler)
	}
	// Report service router path.
	report := router.Group("/api-v1/report")
	{
		report.POST("/add", instance.ReportHandler.AddHandler)
		report.GET("/get-total", instance.ReportHandler.GetTotalHandler)
		report.GET("/list", instance.ReportHandler.ListHandler)
		report.GET("/get-one", instance.ReportHandler.GetOneHandler)
		report.GET("/is-exist", instance.ReportHandler.IsExistHandler)
		report.DELETE("/delete", instance.ReportHandler.DeleteHandler)
		report.PUT("/update", instance.ReportHandler.UpdateHandler)
	}

	// Run on the specified port.
	router.Run(viper.GetString("port"))
}
