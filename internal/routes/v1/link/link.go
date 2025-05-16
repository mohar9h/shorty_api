package link

import (
	"github.com/gin-gonic/gin"
	"shorty_api/internal/database"
	"shorty_api/internal/link"
	"shorty_api/internal/middleware"
)

func RegisterRoutes(router *gin.RouterGroup) {
	routeGroup := router.Group("/link")

	db := database.GetDB()

	linkRepository := link.NewLinkRepository(db)
	linkUsecase := link.NewUsecase(linkRepository)
	linkHandler := link.NewHandler(linkUsecase)

	routeGroup.POST("/", middleware.RateLimiter(), linkHandler.Shorten)
	routeGroup.GET("/:code", linkHandler.Redirect)
}
