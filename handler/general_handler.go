package handler

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	GetTotalHandler(c *gin.Context)
	ListHandler(c *gin.Context)
	AddHandler(c *gin.Context)
	GetOneHandler(c *gin.Context)
	IsExistHandler(c *gin.Context)
	DeleteHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
}
