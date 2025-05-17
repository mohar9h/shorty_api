package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", SubmitPage)
	router.GET("/web/:code", ResultPage)
}
