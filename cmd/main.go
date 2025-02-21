package main

import (
	"discount-service/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api.SetupRoutes(r)

	r.Run()
}
