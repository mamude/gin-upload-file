package server

import (
	"math"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"example.com/mamude/internal/helpers"
	"example.com/mamude/internal/service"
	"example.com/mamude/internal/types"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.LoadHTMLGlob("cmd/web/templates/*")

	r.Static("/assets", "./cmd/web/assets")

	r.GET("/health", s.healthHandler)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})
	r.POST("/send_file", s.sendFileHandler)

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.DB.Health())
}

func (s *Server) sendFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	if !helpers.ValidateFile(file) {
		helpers.BadRequestForFile(c, "arquivo inválido!")
		return
	}

	// upload the file
	fileName := "tmp/" + filepath.Base(uuid.New().String()+".txt")
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		helpers.BadRequest(c, err)
		return
	}

	// counter
	start := time.Now()
	// handle file
	customers := service.SanitizeData(fileName)
	// save to database
	records := s.DB.SaveCustomers(c, customers)
	// send response
	seconds := time.Since(start).Seconds()
	seconds = math.Round(seconds*100) / 100
	data := types.Data{Seconds: seconds, Records: records}
	c.HTML(http.StatusCreated, "result.html", gin.H{"data": data})
}
