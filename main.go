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
		caseInstance.POST("/add", instance.UserHandler.AddHandler)
		caseInstance.GET("/get-total", instance.UserHandler.GetTotalHandler)
		caseInstance.GET("/list", instance.UserHandler.ListHandler)
		caseInstance.GET("/get-one", instance.UserHandler.GetOneHandler)
		caseInstance.GET("/is-exist", instance.UserHandler.IsExistHandler)
		caseInstance.DELETE("/delete", instance.UserHandler.DeleteHandler)
		caseInstance.PUT("/update", instance.UserHandler.UpdateHandler)
	}
	// Static analysis service router path.
	staticAnalysis := router.Group("/api-v1/static-analysis")
	{
		staticAnalysis.POST("/add", instance.UserHandler.AddHandler)
		staticAnalysis.GET("/get-total", instance.UserHandler.GetTotalHandler)
		staticAnalysis.GET("/list", instance.UserHandler.ListHandler)
		staticAnalysis.GET("/get-one", instance.UserHandler.GetOneHandler)
		staticAnalysis.GET("/is-exist", instance.UserHandler.IsExistHandler)
		staticAnalysis.DELETE("/delete", instance.UserHandler.DeleteHandler)
		staticAnalysis.PUT("/update", instance.UserHandler.UpdateHandler)
	}
	// Dynamic analysis service router path.
	dynamicAnalysis := router.Group("/api-v1/dynamicAnalysis")
	{
		dynamicAnalysis.POST("/add", instance.UserHandler.AddHandler)
		dynamicAnalysis.GET("/get-total", instance.UserHandler.GetTotalHandler)
		dynamicAnalysis.GET("/list", instance.UserHandler.ListHandler)
		dynamicAnalysis.GET("/get-one", instance.UserHandler.GetOneHandler)
		dynamicAnalysis.GET("/is-exist", instance.UserHandler.IsExistHandler)
		dynamicAnalysis.DELETE("/delete", instance.UserHandler.DeleteHandler)
		dynamicAnalysis.PUT("/update", instance.UserHandler.UpdateHandler)
	}
	// Video service router path.
	video := router.Group("/api-v1/video")
	{
		video.POST("/add", instance.UserHandler.AddHandler)
		video.GET("/get-total", instance.UserHandler.GetTotalHandler)
		video.GET("/list", instance.UserHandler.ListHandler)
		video.GET("/get-one", instance.UserHandler.GetOneHandler)
		video.GET("/is-exist", instance.UserHandler.IsExistHandler)
		video.DELETE("/delete", instance.UserHandler.DeleteHandler)
		video.PUT("/update", instance.UserHandler.UpdateHandler)
	}
	// Report service router path.
	report := router.Group("/api-v1/report")
	{
		report.POST("/add", instance.UserHandler.AddHandler)
		report.GET("/get-total", instance.UserHandler.GetTotalHandler)
		report.GET("/list", instance.UserHandler.ListHandler)
		report.GET("/get-one", instance.UserHandler.GetOneHandler)
		report.GET("/is-exist", instance.UserHandler.IsExistHandler)
		report.DELETE("/delete", instance.UserHandler.DeleteHandler)
		report.PUT("/update", instance.UserHandler.UpdateHandler)
	}

	// Run on the specified port.
	router.Run(viper.GetString("port"))
}
