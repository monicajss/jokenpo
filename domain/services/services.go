package services

import (
	"jokenpo/domain/entities"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStringResult(input int) string {
	switch input {
	case 1:
		return entities.ROCK
	case 2:
		return entities.PAPER
	case 3:
		return entities.SISSORS
	default:
		return ""
	}
}

func GetMove(c *gin.Context) (entities.Jokenpo, error) {
	var newJokenpo entities.Jokenpo

	if err := c.BindJSON(&newJokenpo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newJokenpo)
		return entities.Jokenpo{}, err
	}

	c.IndentedJSON(http.StatusCreated, newJokenpo)
	return newJokenpo, nil
}

func GeneratePCMoves() string {
	intValue := generateRandomNumber()

	return getStringResult(intValue)
}

func generateRandomNumber() int {
	return rand.Intn(4-1) - 1
}

func CheckResult(c *gin.Context, userMove, pcMove string) {
	if userMove != entities.ROCK && userMove != entities.PAPER && userMove != entities.SISSORS {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "inform a valid move"})
		return
	}

	if userMove == pcMove {
		c.IndentedJSON(http.StatusOK, gin.H{"result": "tie"})
		return
	}

	if userMove == entities.ROCK && pcMove == entities.SISSORS || userMove == entities.PAPER && pcMove == entities.ROCK || userMove == entities.SISSORS && pcMove == entities.PAPER {
		c.IndentedJSON(http.StatusOK, gin.H{"result": "victory"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": "defeat"})
}
