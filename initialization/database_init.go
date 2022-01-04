package initialization

import (
	"fmt"
	"log"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (builder *InitializationList) InitializeDatabase() *InitializationList {
	// Get the database source name.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.name"),
		viper.GetString("mysql.charset"))

	// To recover the connection error.
	var err error
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	builder.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	builder.DB.AutoMigrate(&model.User{})
	builder.DB.AutoMigrate(&model.Case{})
	builder.DB.AutoMigrate(&model.StaticAnalysis{})
	builder.DB.AutoMigrate(&model.DynamicAnalysis{})
	builder.DB.AutoMigrate(&model.Video{})
	builder.DB.AutoMigrate(&model.Report{})
	log.Println("Data successfully initialized.")

	return builder
}
