package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Teste(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Teste ok",
	})
}
