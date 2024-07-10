package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context, msg string) {
	c.JSON(http.StatusCreated, gin.H{"message": msg})
}

func BadRequest(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
