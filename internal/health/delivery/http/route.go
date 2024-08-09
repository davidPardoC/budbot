package http

import "github.com/gin-gonic/gin"

type WebhookRouter struct {
	router *gin.Engine
}

func NewHealthRouter(router *gin.Engine) *WebhookRouter {
	return &WebhookRouter{router: router}
}

func (r *WebhookRouter) SetupHelthRouter() {
	r.router.GET("/health", HealtCheckHandler)
}
