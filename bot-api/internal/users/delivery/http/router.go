package http

import (
	budgetRepository "github.com/davidPardoC/budbot/internal/budgets/repository"
	transactionsRepository "github.com/davidPardoC/budbot/internal/transactions/repository"
	"gorm.io/gorm"

	"github.com/davidPardoC/budbot/internal/users/repository"
	"github.com/davidPardoC/budbot/internal/users/usecases"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	gin      *gin.Engine
	handlers *UserHandlers
}

func NewUserRouter(gin *gin.Engine, db *gorm.DB) *UserRouter {
	userRepository := repository.NewUserRepository(db)
	budgetRepo := budgetRepository.NewBudgetRepository(db)
	transactionsRepo := transactionsRepository.NewTransactionsRepository(db)
	userUseCases := usecases.NewUserUsecases(userRepository, budgetRepo, transactionsRepo)
	handlers := NewUserHandler(userUseCases)
	return &UserRouter{handlers: handlers, gin: gin}
}

func (r *UserRouter) SetupRoutes() {
	users := r.gin.Group("/api/v1/users")
	{
		users.GET("/:user_id/stats", r.handlers.GetStats)
		users.GET("/:user_id/transactions", r.handlers.GetTransactions)
		users.GET("/:user_id/transactions-grouped", r.handlers.GetTransactionsGrouped)
	}
}
