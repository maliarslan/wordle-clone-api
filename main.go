package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maliarslan/wordle-clone-api/internal/routes"
)

func main() {
	router := gin.Default()

	routes.HealthRouter(router)
	routes.WordleRouter(router)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}
