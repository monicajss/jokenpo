package main

import (
	"jokenpo/domain/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/move", controllers.HandleMove)

	router.Run("0.0.0.0:8000")
}
