package http

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type FrontendRouter struct {
	gin *gin.Engine
}

func NewFrontendRouter(gin *gin.Engine) *FrontendRouter {
	return &FrontendRouter{gin: gin}
}

func (r *FrontendRouter) SetupFrontendRouter() {
	r.gin.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))
}
