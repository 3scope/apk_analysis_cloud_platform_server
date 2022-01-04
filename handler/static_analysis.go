package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/sanscope/apk_analysis_cloud_platform_server/enum"
	_ "github.com/sanscope/apk_analysis_cloud_platform_server/modify_log"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"

	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/service"
)

type StaticAnalysisHandler struct {
	Srv *service.StaticAnalysisService
}

//
type StaticAnalysisHandlerInterface interface {
	GetTotalHandler(c *gin.Context)
	ListHandler(c *gin.Context)
	AddHandler(c *gin.Context)
	GetOneHandler(c *gin.Context)
	IsExistHandler(c *gin.Context)
	DeleteHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
}

func (h *StaticAnalysisHandler) GetTotalHandler(c *gin.Context) {
	var entity repository.StaticAnalysisEntity
	request := RequestFactory(c, entity)
	if request == nil {
		log.Println(errors.New("request is empty"))
		return
	}

	if result, err := h.Srv.GetTotal(request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusInternalServerError, res)

		log.Println(err)
		return
	} else {
		// If successful, pass in the result value.
		res := ResponseFactory(enum.SuccessfulOpration, result)
		c.JSON(http.StatusOK, res)

		log.Println(res)
		return
	}
}

func (h *StaticAnalysisHandler) ListHandler(c *gin.Context) {
	var entity repository.StaticAnalysisEntity
	request := RequestFactory(c, entity)
	if request == nil {
		log.Println(errors.New("request is empty"))
		return
	}

	if result, err := h.Srv.List(request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusInternalServerError, res)

		log.Println(err)
		return
	} else {
		// If successful, pass in the result value.
		res := ResponseFactory(enum.SuccessfulOpration, result)
		c.JSON(http.StatusOK, res)

		log.Println(res)
		return
	}
}

func (h *StaticAnalysisHandler) AddHandler(c *gin.Context) {
	var entity repository.StaticAnalysisEntity
	request := RequestFactory(c, entity)
	if request == nil {
		log.Println(errors.New("request is empty"))
		return
	}

	if result, err := h.Srv.Add(request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusInternalServerError, res)

		log.Println(err)
		return
	} else {
		res := ResponseFactory(enum.SuccessfulOpration, result)
		c.JSON(http.StatusOK, res)

		log.Println(res)
		return
	}
}

func (h *StaticAnalysisHandler) GetOneHandler(c *gin.Context) {
	var entity repository.StaticAnalysisEntity
	request := RequestFactory(c, entity)
	if request == nil {
		log.Println(errors.New("request is empty"))
		return
	}

	if result, err := h.Srv.GetOne(request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusInternalServerError, res)

		log.Println(err)
		return
	} else {
		res := ResponseFactory(enum.SuccessfulOpration, result)
		c.JSON(http.StatusOK, res)

		log.Println(res)
		return
	}
}

func (h *StaticAnalysisHandler) IsExistHandler(c *gin.Context) {
	var entity repository.StaticAnalysisEntity
	request := RequestFactory(c, entity)
	if request == nil {
		log.Println(errors.New("request is empty"))
		return
	}

	// Return boolean value.
	if result, err := h.Srv.IsExist(request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusInternalServerError, res)

		log.Println(err)
		return
	} else {
		res := ResponseFactory(enum.SuccessfulOpration, result)
		c.JSON(http.StatusOK, res)

		log.Println(res)
		return
	}
}

func (h *StaticAnalysisHandler) DeleteHandler(c *gin.Context) {
	var entity repository.StaticAnalysisEntity
	request := RequestFactory(c, entity)
	if request == nil {
		log.Println(errors.New("request is empty"))
		return
	}

	// Server internal error.
	if result, err := h.Srv.Delete(request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusInternalServerError, res)

		log.Println(err)
		return
	} else {
		res := ResponseFactory(enum.SuccessfulOpration, result)
		c.JSON(http.StatusOK, res)

		log.Println(res)
		return
	}
}

func (h *StaticAnalysisHandler) UpdateHandler(c *gin.Context) {
	var entity repository.StaticAnalysisEntity
	request := RequestFactory(c, entity)
	if request == nil {
		log.Println(errors.New("request is empty"))
		return
	}

	// Server internal error.
	if result, err := h.Srv.Update(request); err != nil {
		res := ResponseFactory(enum.FailedOpration, nil)
		c.JSON(http.StatusInternalServerError, res)

		log.Println(err)
		return
	} else {
		res := ResponseFactory(enum.SuccessfulOpration, result)
		c.JSON(http.StatusOK, res)

		log.Println(res)
		return
	}
}
