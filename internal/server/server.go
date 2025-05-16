package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"shorty_api/internal/config"
	"shorty_api/internal/middleware"
	"shorty_api/internal/routes"
	"shorty_api/internal/validation"
)

func SetupServer(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.NoRoute(middleware.NotFoundHandler())
	router.NoMethod(middleware.MethodNotAllowedHandler())

	if err := validation.InitValidator(); err != nil {
		log.Fatal(err)
	}
	router.Use(middleware.DefaultStructuredLogger(cfg))
	router.Use(gin.Recovery(), gin.Logger()) //, middleware.RateLimiter())

	routes.RegisterRoutes(router)

	return router
}
