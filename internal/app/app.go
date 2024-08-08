package app

import (
	"fmt"

	"github.com/davidPardoC/budbot/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db  *gorm.DB
	gin *gin.Engine
	cfg config.Config
}

func NewApp(db *gorm.DB, gin *gin.Engine, cfg config.Config) *App {
	return &App{
		db:  db,
		gin: gin,
		cfg: cfg,
	}
}

func (a *App) Run() {
	a.gin.Run(fmt.Sprintf(":%s", a.cfg.Server.Port))
}
