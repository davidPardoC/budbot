package main

import (
	"time"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/database"
	"github.com/davidPardoC/budbot/internal/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	postgresDatabase := database.Connect(config)

	database.Migrate(postgresDatabase)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE", "PATCH"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour

	router.Use(cors.New(corsConfig))

	app.NewApp(postgresDatabase, router, config).Run()

}
