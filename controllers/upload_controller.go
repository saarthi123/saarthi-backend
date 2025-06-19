package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No file is received",
		})
		return
	}

	// Create uploads folder if not exists
	uploadPath := "uploads"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}

	// Save file with a unique name
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	savePath := filepath.Join(uploadPath, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to save the file",
		})
		return
	}

	// Return the URL (assuming it's served via static files)
	c.JSON(http.StatusOK, gin.H{
		"url": fmt.Sprintf("http://localhost:3000/uploads/%s", filename),
	})
}
