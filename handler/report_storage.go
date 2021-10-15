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

type ReportHandler struct {
	Srv *service.ReportService
}

func (h *ReportHandler) GetTotalHandler(c *gin.Context) {
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

func (h *ReportHandler) ListHandler(c *gin.Context) {
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

func (h *ReportHandler) AddHandler(c *gin.Context) {
	m := model.ReportStorage{}
	// Client request error.
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Add(m)
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

func (h *ReportHandler) GetHandler(c *gin.Context) {
	m := model.ReportStorage{}
	// Client request error.
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	value := h.Srv.Get(m)
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *ReportHandler) IsExistHandler(c *gin.Context) {
	m := model.ReportStorage{}
	// Client request error.
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Return boolean value.
	value := h.Srv.IsExist(m)
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *ReportHandler) DeleteHandler(c *gin.Context) {
	m := model.ReportStorage{}
	// Client request error.
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Delete(m)
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

func (h *ReportHandler) UpdateHandler(c *gin.Context) {
	m := model.ReportStorage{}
	// Client request error.
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Update(m)
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
