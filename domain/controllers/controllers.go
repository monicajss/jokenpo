package controllers

import (
	"jokenpo/domain/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleMove(c *gin.Context) {
	userMove, err := services.GetMove(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	pcMove := services.GeneratePCMoves()
	services.CheckResult(c, strings.ToLower(userMove.Move), pcMove)
}
