package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Gin basic settings.
	router := gin.Default()
	router.Use(CORSSettings())
	gin.SetMode(viper.GetString("mode"))

	// User service router path.
	user := router.Group("/api-v1/user")
	{
		user.POST("/add", UserHandler.AddHandler)
		user.GET("/get-total", UserHandler.GetTotalHandler)
		user.GET("/list", UserHandler.ListHandler)
		user.GET("/get", UserHandler.GetHandler)
		user.GET("/is-exist", UserHandler.IsExistHandler)
		user.DELETE("/delete", UserHandler.DeleteHandler)
		user.PUT("/update", UserHandler.UpdateHandler)
	}
	// Apk service router path.
	apk := router.Group("/api-v1/apk")
	{
		apk.POST("/add", ApkHandler.AddHandler)
		apk.GET("/get-total", ApkHandler.GetTotalHandler)
		apk.GET("/list", ApkHandler.ListHandler)
		apk.GET("/get", ApkHandler.GetHandler)
		apk.GET("/is-exist", ApkHandler.IsExistHandler)
		apk.DELETE("/delete", ApkHandler.DeleteHandler)
		apk.PUT("/update", ApkHandler.UpdateHandler)
	}
	// Video service router path.
	video := router.Group("/api-v1/video")
	{
		video.POST("/add", VideoHandler.AddHandler)
		video.GET("/get-total", VideoHandler.GetTotalHandler)
		video.GET("/list", VideoHandler.ListHandler)
		video.GET("/get", VideoHandler.GetHandler)
		video.GET("/is-exist", VideoHandler.IsExistHandler)
		video.DELETE("/delete", VideoHandler.DeleteHandler)
		video.PUT("/update", VideoHandler.UpdateHandler)
	}
	// Report service router path.
	report := router.Group("/api-v1/report")
	{
		report.POST("/add", ReportHandler.AddHandler)
		report.GET("/get-total", UserHandler.GetTotalHandler)
		report.GET("/list", UserHandler.ListHandler)
		report.GET("/get", UserHandler.GetHandler)
		report.GET("/is-exist", UserHandler.IsExistHandler)
		report.DELETE("/delete", UserHandler.DeleteHandler)
		report.PUT("/update", UserHandler.UpdateHandler)
	}
	// Run on the specified port.
	router.Run(viper.GetString("port"))
}
