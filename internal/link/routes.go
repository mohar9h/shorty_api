package link

import (
	"github.com/gin-gonic/gin"
	"shorty_api/internal/middleware"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("/", middleware.RateLimiter(), handler.Shorten)
	router.GET("/:code", handler.Redirect)
}
