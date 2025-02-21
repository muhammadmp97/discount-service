package api

import (
	"discount-service/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
}
