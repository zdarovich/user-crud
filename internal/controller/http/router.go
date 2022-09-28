package http

import (
	"binance-order-matcher/internal/service"
	"binance-order-matcher/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(handler *gin.Engine, l logger.Interface, t service.UserRepo) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Routers
	h := handler.Group("/v1")
	{
		newUserRoutes(h, t, l)
	}
}
