package token

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sanscope/apk_analysis_cloud_platform_server/enum"
)

var secretKey = []byte("apk_analysis_cloud_platform")

type JWTClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

// To get json web token.
func generateToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func verifyToken(signedToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("assertion failed")
	}
	// To verify token that token has expired.
	err = token.Claims.Valid()
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// The refresh handler
func RefreshToken(c *gin.Context) {
	// The Header is essentially a map[string]string.
	tokenHeader := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(tokenHeader, " ", 2)
	// Authorization: Bearer <token>
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    enum.UnauthorizedOpration,
			"message": enum.UnauthorizedOpration.String(),
			"data":    nil,
		})
		return
	}

	claims, err := verifyToken(parts[1])
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    enum.UnauthorizedOpration,
			"message": enum.UnauthorizedOpration.String(),
			"data":    nil,
		})
		return
	}
	// Reset the expiration time of the token.
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = claims.IssuedAt + 60*60
	signedToken, err := generateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    enum.FailedOpration,
			"message": enum.FailedOpration.String(),
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
