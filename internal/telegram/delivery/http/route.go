package http

import "github.com/gin-gonic/gin"

type WebhookRouter struct {
	router *gin.Engine
}

func NewWebhookRouter(router *gin.Engine) *WebhookRouter {
	return &WebhookRouter{router: router}
}

func (r *WebhookRouter) SetupWebhookRouter() {
	r.router.POST("/api/v1/webhook/telegram", WebHookHandler)
}
