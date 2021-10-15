package handler

import (
	"log"
	"net/http"

	"github.com/sanscope/apk_analysis_cloud_platform_server/enum"
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	_ "github.com/sanscope/apk_analysis_cloud_platform_server/modify_log"

	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
	"github.com/sanscope/apk_analysis_cloud_platform_server/service"
)

type ApkHandler struct {
	Srv *service.ApkService
}

func (h *ApkHandler) GetTotalHandler(c *gin.Context) {
	var request repository.Query
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	value := h.Srv.GetTotal(&request)
	// Same structure as Response struct.
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *ApkHandler) ListHandler(c *gin.Context) {
	var request repository.Query
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	value := h.Srv.List(&request)
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *ApkHandler) AddHandler(c *gin.Context) {
	a := model.ApkStorage{}
	// Client request error.
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Add(a)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    enum.FailedOpration,
			"message": message,
			"data":    nil,
		})
		log.Panicln(message)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": message,
		"data":    nil,
	})
}

func (h *ApkHandler) GetHandler(c *gin.Context) {
	a := model.ApkStorage{}
	// Client request error.
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	value := h.Srv.Get(a)
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *ApkHandler) IsExistHandler(c *gin.Context) {
	a := model.ApkStorage{}
	// Client request error.
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Return boolean value.
	value := h.Srv.IsExist(a)
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *ApkHandler) DeleteHandler(c *gin.Context) {
	a := model.ApkStorage{}
	// Client request error.
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Delete(a)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    enum.FailedOpration,
			"message": message,
			"data":    nil,
		})
		log.Panicln(message)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": message,
		"data":    nil,
	})
}

func (h *ApkHandler) UpdateHandler(c *gin.Context) {
	a := model.ApkStorage{}
	// Client request error.
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Update(a)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    enum.FailedOpration,
			"message": message,
			"data":    nil,
		})
		log.Panicln(message)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": message,
		"data":    nil,
	})
}
