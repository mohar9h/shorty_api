package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"shorty_api/internal/config"
	"shorty_api/internal/database"
	"shorty_api/internal/middleware"
	"shorty_api/internal/routes"
	"shorty_api/internal/validation"
)

func SetupServer(cfg *config.Config) *gin.Engine {

	router := gin.New()

	router.Use(
		middleware.DefaultStructuredLogger(cfg),
		gin.Recovery(),
		gin.Logger(),
	)

	router.NoRoute(middleware.NotFoundHandler())
	router.NoMethod(middleware.MethodNotAllowedHandler())

	if err := validation.InitValidator(); err != nil {
		log.Fatalf("Validator init failed: %v", err)
	}

	db := database.GetDB()

	routes.Register(router, db)

	return router
}
