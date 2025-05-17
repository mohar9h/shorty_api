package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"shorty_api/internal/web"

	linkModule "shorty_api/internal/link"
)

func Register(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/")

	registerLinkRoutes(routes, db)
	registerWebpageRoutes(routes)
}

func registerLinkRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := linkModule.NewLinkRepository(db)
	usecase := linkModule.NewUsecase(repo)
	handler := linkModule.NewHandler(usecase)

	linkModule.RegisterRoutes(router.Group("/link"), handler)
}

func registerWebpageRoutes(router *gin.RouterGroup) {
	web.RegisterRoutes(router.Group("/"))
}
