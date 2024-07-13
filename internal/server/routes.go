package server

import (
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"example.com/mamude/internal/helpers"
	"example.com/mamude/internal/repository"
	"example.com/mamude/internal/service"
	"example.com/mamude/internal/types"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.LoadHTMLGlob(os.Getenv("TEMPLATE"))
	r.Static("/assets", os.Getenv("ASSETS"))

	r.GET("/ping", s.pingHandler)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})
	r.POST("/send_file", s.sendFileHandler)

	return r
}

func (s *Server) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (s *Server) sendFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		helpers.BadRequestForFile(c, "selecione um arquivo!")
		return
	}

	if !helpers.ValidateFile(file) {
		helpers.BadRequestForFile(c, "arquivo inválido!")
		return
	}

	// upload the file
	tempFiles := os.Getenv("TEMP_FILES")
	fileName := tempFiles + filepath.Base(uuid.New().String()+".txt")
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		helpers.BadRequest(c, err)
		return
	}

	// counter
	start := time.Now()
	// handle file
	customers := service.SanitizeData(fileName)
	// save to database
	repo := repository.NewCustomerRepository(s.DB)
	records := repo.SaveData(c, customers)
	// send response
	seconds := time.Since(start).Seconds()
	seconds = math.Round(seconds*100) / 100
	data := types.Data{Seconds: seconds, Records: records}
	c.HTML(http.StatusCreated, "result.html", gin.H{"data": data})
}
