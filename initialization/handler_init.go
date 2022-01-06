package initialization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanscope/apk_analysis_cloud_platform_server/handler"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
	"github.com/sanscope/apk_analysis_cloud_platform_server/service"
)

func (builder *InitializationList) InitializeHandler() *InitializationList {
	// Interface is used for abstraction.
	builder.UserHandler = &handler.UserHandler{
		Srv: &service.UserService{
			Repository: &repository.UserRepository{
				DB: builder.DB,
			},
		},
	}
	builder.CaseHandler = &handler.CaseHandler{
		Srv: &service.CaseService{
			Repository: &repository.CaseRepository{
				DB: builder.DB,
			},
		},
	}
	builder.StaticAnalysisHandler = &handler.StaticAnalysisHandler{
		Srv: &service.StaticAnalysisService{
			Repository: &repository.StaticAnalysisRepository{
				DB: builder.DB,
			},
		},
	}
	builder.DynamicAnalysisHandler = &handler.DynamicAnalysisHandler{
		Srv: &service.DynamicAnalysisService{
			Repository: &repository.DynamicAnalysisRepository{
				DB: builder.DB,
			},
		},
	}
	builder.ReportHandler = &handler.ReportHandler{
		Srv: &service.ReportService{
			Repository: &repository.ReportRepository{
				DB: builder.DB,
			},
		},
	}
	builder.VideoHandler = &handler.VideoHandler{
		Srv: &service.VideoService{
			Repository: &repository.VideoRepository{
				DB: builder.DB,
			},
		},
	}

	return builder
}

func CORSSettings() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// Handling OPTIONS requests generated across domains.
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
