package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"example.com/mamude/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.MaxMultipartMemory = 6 << 20

	r.GET("/health", s.healthHandler)

	r.Static("/assets", "./cmd/web/assets")

	r.GET("/", func(c *gin.Context) {
		templ.Handler(web.UploadFormFile()).ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/send_file", s.sendFileHandler)

	r.POST("/hello", func(c *gin.Context) {
		web.HelloWebHandler(c.Writer, c.Request)
	})

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) sendFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Campo requirido: %s": err.Error()})
		return
	}

	// upload the file
	fileName := "tmp/" + filepath.Base(uuid.New().String()+".txt")
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Arquivo inválido: %s": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": fmt.Sprintf("%s processado!", file.Filename)})
}
