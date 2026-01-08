package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maliarslan/wordle-clone-api/internal/utils"
)

type Wordle struct {
	Wordle string `json:"wordle"`
}

type CheckWordle struct {
	Correct bool `json:"correct"`
}

func GetWordle(c *gin.Context) {
	word, err := utils.GetRandomWord(5)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, word)
}

func PostWordle(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "wordle")
}
