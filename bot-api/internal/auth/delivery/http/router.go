package http

import (
	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/auth/usecases"
	userRepository "github.com/davidPardoC/budbot/internal/users/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authRouter struct {
	gin      *gin.Engine
	handlers *AuthHandlers
}

func NewAuthRouter(gin *gin.Engine, db *gorm.DB, config config.Config) *authRouter {

	userRepository := userRepository.NewUserRepository(db)
	useCases := usecases.NewAuthUseCasesImp(userRepository, config)
	handlers := NewAuthHandlers(useCases, config)

	return &authRouter{gin: gin, handlers: handlers}
}

func (r *authRouter) SetupRoutes() {

	group := r.gin.Group("/api/v1/auth")
	{
		group.GET("/telegram/callback", r.handlers.oAuthTelegramHandler)
	}
}
