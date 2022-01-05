package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/enum"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

// The entity type will change dynamically.
func RequestFactory(c *gin.Context, entityType enum.EntityType) *repository.Request {
	var request repository.Request

	// If use json to pass data, this method will be changed to "c.ShouldBindBodyWith".
	if err := c.ShouldBind(&request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusBadRequest, res)

		log.Println(err)
		return nil
	}

	if entity, err := EntityFactory(c, entityType); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusBadRequest, res)

		log.Println(err)
		return nil
	} else {
		// Assign value if binding is successful.
		// The entity variable is struct, but the "Entity" attribute is pointer.
		request.Entity = entity
	}

	return &request
}

func EntityFactory(c *gin.Context, entityType enum.EntityType) (interface{}, error) {
	var entity interface{}
	var err error
	switch entityType {
	case enum.UserEntity:
		ins := repository.UserEntity{}
		err = c.ShouldBind(&ins)
		// The entity is the pointer.
		entity = &ins
	case enum.CaseEntity:
		ins := repository.CaseEntity{}
		err = c.ShouldBind(&ins)
		// The entity is the pointer.
		entity = &ins
	case enum.StaticAnalysisEntity:
		ins := repository.StaticAnalysisEntity{}
		err = c.ShouldBind(&ins)
		// The entity is the pointer.
		entity = &ins
	case enum.DynamicAnalysisEntity:
		ins := repository.DynamicAnalysisEntity{}
		err = c.ShouldBind(&ins)
		// The entity is the pointer.
		entity = &ins
	case enum.ReportEntity:
		ins := repository.ReportEntity{}
		err = c.ShouldBind(&ins)
		// The entity is the pointer.
		entity = &ins
	case enum.VideoEntity:
		ins := repository.VideoEntity{}
		err = c.ShouldBind(&ins)
		// The entity is the pointer.
		entity = &ins
	}
	return entity, err
}
