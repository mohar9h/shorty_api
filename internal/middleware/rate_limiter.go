package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorty_api/internal/config"
	response "shorty_api/internal/utils"
)

func RateLimiter() gin.HandlerFunc {
	cfg := config.GetConfig()
	limiter := tollbooth.NewLimiter(float64(cfg.RateLimit.RequestsPerSecond), nil)
	limiter.SetBurst(cfg.RateLimit.BurstSize)

	return func(context *gin.Context) {
		err := tollbooth.LimitByRequest(limiter, context.Writer, context.Request)
		if err != nil {
			_ = response.Error(context.Writer, http.StatusTooManyRequests, "Too many requests.", err.Error())
			context.Abort()
			return
		}
		context.Next()
	}
}
