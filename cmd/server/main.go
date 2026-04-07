package main

import (
	"log"
	"tidy/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	r.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server running in" + cfg.AppEnv,
		})
	})
	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
