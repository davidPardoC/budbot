package http

import "github.com/gin-gonic/gin"

type authRouter struct {
	gin *gin.Engine
}

func NewAuthRouter(gin *gin.Engine) *authRouter {
	return &authRouter{gin: gin}
}

func (r *authRouter) SetupRoutes() {
	group := r.gin.Group("/api/v1/auth")
	{
		group.GET("/telegram/callback", oAuthTelegramHandler)
	}
}
