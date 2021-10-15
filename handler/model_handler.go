package handler

import "github.com/gin-gonic/gin"

type ModelHandlerInterface interface {
	GetTotalHandler(c *gin.Context)
	ListHandler(c *gin.Context)
	AddHandler(c *gin.Context)
	GetHandler(c *gin.Context)
	IsExistHandler(c *gin.Context)
	DeleteHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
}
