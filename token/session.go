package token

import (
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/service"
)

type SessionHandler struct {
	Srv *service.UserService
}

func SessionInit(router *gin.Engine) error {
	// To set cache into redis.
	redis, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// The options of cookie.
	redis.Options(sessions.Options{
		Path:   "/",
		Domain: "",
		MaxAge: 600,
	})
	if err != nil {
		return err
	}
	router.Use(sessions.Sessions("redis", redis))
	return nil
}

func (sh *SessionHandler) Login(c *gin.Context) {
	redis := sessions.Default(c)
	// TODO: Use "ShouldBind" to bind request.
	redis.Set("userID", time.Now())
	redis.Get("userID")
	redis.Delete("userID")
	redis.Save()
}
