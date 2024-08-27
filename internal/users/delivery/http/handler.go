package http

import (
	"strconv"

	"github.com/davidPardoC/budbot/internal/users/usecases"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	userUseCases usecases.IUserUseCases
}

type RouteParams struct {
	UserId string `uri:"user_id" binding:"required"`
}

type RouteQuery struct {
	Month string `form:"month" binding:"required"`
	Year  string `form:"year" binding:"required"`
}

func NewUserHandler(userUseCases usecases.IUserUseCases) *UserHandlers {
	return &UserHandlers{userUseCases: userUseCases}
}

func (u *UserHandlers) GetStats(c *gin.Context) {
	var params RouteParams
	var query RouteQuery

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	parsedUserId, err := strconv.ParseInt(params.UserId, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	parsedMonth, err := strconv.ParseInt(query.Month, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	parsedYear, err := strconv.ParseInt(query.Year, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	stats, _ := u.userUseCases.GetStatsBetweenDates(parsedUserId, int(parsedMonth), int(parsedYear))

	c.JSON(200, stats)
}
