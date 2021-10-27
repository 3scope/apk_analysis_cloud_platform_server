package token

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sanscope/apk_analysis_cloud_platform_server/enum"
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/service"
)

type UserActionInterface interface {
	// POST method.
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
}

type UserActionHandler struct {
	Srv *service.UserService
}

func (h *UserActionHandler) Login(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
			"data":    nil,
		})
		return
	}
	// Instantiate custom claims.
	claims := &JWTClaims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			// Effective time.
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60,
			IssuedAt:  time.Now().Unix(),
		},
	}
	signedToken, err := generateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    enum.BusyStatus,
			"message": enum.BusyStatus.String(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    enum.SuccessfulOpration,
		"message": enum.SuccessfulOpration.String(),
		"data":    signedToken,
	})
}

func (h *UserActionHandler) Register(c *gin.Context) {

}

func (h *UserActionHandler) Logout(c *gin.Context) {

}
