package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/sanscope/apk_analysis_cloud_platform_server/enum"
	_ "github.com/sanscope/apk_analysis_cloud_platform_server/modify_log"

	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/service"
)

type VideoHandler struct {
	Srv *service.VideoService
}

//
type VideoHandlerInterface interface {
	GetTotalHandler(c *gin.Context)
	ListHandler(c *gin.Context)
	AddHandler(c *gin.Context)
	GetOneHandler(c *gin.Context)
	IsExistHandler(c *gin.Context)
	DeleteHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
}

func (h *VideoHandler) GetTotalHandler(c *gin.Context) {
	request := RequestFactory(c, enum.VideoEntity)
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

func (h *VideoHandler) ListHandler(c *gin.Context) {
	request := RequestFactory(c, enum.VideoEntity)
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

func (h *VideoHandler) AddHandler(c *gin.Context) {
	request := RequestFactory(c, enum.VideoEntity)
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

func (h *VideoHandler) GetOneHandler(c *gin.Context) {
	request := RequestFactory(c, enum.VideoEntity)
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

func (h *VideoHandler) IsExistHandler(c *gin.Context) {
	request := RequestFactory(c, enum.VideoEntity)
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

func (h *VideoHandler) DeleteHandler(c *gin.Context) {
	request := RequestFactory(c, enum.VideoEntity)
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

func (h *VideoHandler) UpdateHandler(c *gin.Context) {
	request := RequestFactory(c, enum.VideoEntity)
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
