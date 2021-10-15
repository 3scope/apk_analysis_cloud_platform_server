package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/sanscope/apk_analysis_cloud_platform_server/config"
	"github.com/sanscope/apk_analysis_cloud_platform_server/handler"
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	_ "github.com/sanscope/apk_analysis_cloud_platform_server/modify_log"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
	"github.com/sanscope/apk_analysis_cloud_platform_server/service"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	UserHandler   *handler.UserHandler
	ApkHandler    *handler.ApkHandler
	VideoHandler  *handler.VideoHandler
	ReportHandler *handler.ReportHandler
)

func initDB() {
	log.Println("Database initialization.")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.name"),
		viper.GetString("database.charset"))
	// To recover the connection error.
	var err error
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.ApkStorage{})
	DB.AutoMigrate(&model.VideoStorage{})
	DB.AutoMigrate(&model.ReportStorage{})
	log.Println("Data successfully initialized.")
}

// func initHandler(*handler.ModelHandlerInterface) {
// 	handler.NewHandler()
// }

func initHandler() {
	// Interface is used for abstraction.
	UserHandler = &handler.UserHandler{
		Srv: &service.UserService{
			Repository: &repository.UserRepository{
				DB: DB,
			},
		},
	}
	ApkHandler = &handler.ApkHandler{
		Srv: &service.ApkService{
			Repository: &repository.ApkRepository{
				DB: DB,
			},
		},
	}
	VideoHandler = &handler.VideoHandler{
		Srv: &service.VideoService{
			Repository: &repository.VideoRepository{
				DB: DB,
			},
		},
	}
	ReportHandler = &handler.ReportHandler{
		Srv: &service.ReportService{
			Repository: &repository.ReportRepository{
				DB: DB,
			},
		},
	}
}

func init() {
	initDB()
	initHandler()
}

func CORSSettings() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// Handling OPTIONS requests generated across domains.
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
