package v1

import (
	"github.com/gin-gonic/gin"
	"shorty_api/internal/routes/v1/link"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	v1Group := routerGroup.Group("/v1")
	link.RegisterRoutes(v1Group)
}
