package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/enum"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

// The entity type will change dynamically.
func RequestFactory(c *gin.Context, entity interface{}) *repository.Request {
	var request repository.Request

	// If use json to pass data, this method will be changed to "c.ShouldBindBodyWith".
	if err := c.ShouldBind(&request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusBadRequest, res)

		log.Println(err)
		return nil
	}

	if err := c.ShouldBind(&entity); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusBadRequest, res)

		log.Println(err)
		return nil
	}

	// Assign value if binding is successful.
	// The entity variable is struct, but the "Entity" attribute is pointer.
	request.Entity = &entity
	return &request
}

// Only bind entiy.
func EntityFactory(c *gin.Context, entity interface{}) interface{} {
	// Client request error.
	if err := c.ShouldBind(&entity); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusBadRequest, res)

		log.Println(err)
		return nil
	}

	return &entity
}
