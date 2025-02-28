package api

import (
	"discount-service/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handler.Ping)

	// Admin endpoints
	r.GET("/discounts", handler.List)
	r.GET("/discounts/:id", handler.Show)
	r.POST("/discounts", handler.Create)
	r.DELETE("/discounts/:id", handler.Delete)
	r.PUT("/discounts/:id", handler.Update)

	// User endpoints
	r.GET("/discounts/:code/verify", handler.Verify)
	r.GET("/discounts/:code/apply", handler.Apply)
	r.GET("/discounts/:code/cancel", handler.Cancel)
}
