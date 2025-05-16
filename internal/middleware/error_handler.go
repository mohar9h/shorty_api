package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	response "shorty_api/internal/utils"
)

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = response.NotFound(c.Writer, "Route not found")
	}
}

func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = response.Error(c.Writer, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}
}
