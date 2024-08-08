package main

import (
	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/database"
	"github.com/davidPardoC/budbot/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	database := database.Connect(config)
	router := gin.Default()
	app.NewApp(database, router, config).Run()
}
