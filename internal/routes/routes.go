package routes

import (
	"github.com/gin-gonic/gin"
	v1 "shorty_api/internal/routes/v1"
)

func RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	v1.RegisterRoutes(apiGroup)
}
