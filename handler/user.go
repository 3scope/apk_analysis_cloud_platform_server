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

type UserHandler struct {
	Srv *service.UserService
}

func (h *UserHandler) GetTotalHandler(c *gin.Context) {
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

func (h *UserHandler) ListHandler(c *gin.Context) {
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

func (h *UserHandler) AddHandler(c *gin.Context) {
	var u model.User
	// Client request error.
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Add(u)
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

func (h *UserHandler) GetHandler(c *gin.Context) {
	u := model.User{}
	// Client request error.
	// Http method has nothing to do with the parameters required by the function.
	if err := c.ShouldBindQuery(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	value := h.Srv.Get(u)
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *UserHandler) IsExistHandler(c *gin.Context) {
	u := model.User{}
	// Client request error.
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Return boolean value.
	value := h.Srv.IsExist(u)
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    value,
	})
}

func (h *UserHandler) DeleteHandler(c *gin.Context) {
	u := model.User{}
	// Client request error.
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Delete(u)
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

func (h *UserHandler) UpdateHandler(c *gin.Context) {
	u := model.User{}
	// Client request error.
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		log.Panicln(err)
		return
	}
	// Server internal error.
	ok, message := h.Srv.Update(u)
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
