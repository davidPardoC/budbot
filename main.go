package main

import (
	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/database"
	"github.com/davidPardoC/budbot/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	postgresDatabase := database.Connect(config)

	if config.Server.Env == "local" {
		database.Migrate(postgresDatabase)
	}

	router := gin.Default()
	app.NewApp(postgresDatabase, router, config).Run()

}
