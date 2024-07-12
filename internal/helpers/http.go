package helpers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context, msg string) {
	c.JSON(http.StatusCreated, gin.H{"message": msg})
}

func BadRequest(c *gin.Context, err error) {
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
	}
}

func BadRequestForFile(c *gin.Context, message string) {
	c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": message})
}

func ValidateFile(file *multipart.FileHeader) bool {
	headers := file.Header.Get("Content-Type")
	return headers == "text/plain"
}
