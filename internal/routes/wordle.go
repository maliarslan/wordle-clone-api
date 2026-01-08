package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maliarslan/wordle-clone-api/internal/handlers"
)

func WordleRouter(router *gin.Engine) {
	wordle := router.Group("/wordle")
	{
		wordle.GET("", handlers.GetWordle)
		wordle.POST("/check", handlers.PostWordle)
	}
}
