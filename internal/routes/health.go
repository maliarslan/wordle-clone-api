package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maliarslan/wordle-clone-api/internal/handlers"
)

func HealthRouter(router *gin.Engine) {
	health := router.Group("/health")
	{
		health.GET("", handlers.HealthCheck)
	}
}
