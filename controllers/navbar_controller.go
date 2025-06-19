package controllers

import (
	"net/http"
	"github.com/saarthi123/saarthi-backend/models"
  "github.com/gin-gonic/gin")

func GetCommands(c *gin.Context) {
	var commands []models.Command
	models.DB.Find(&commands)
  c.JSON(http.StatusOK, gin.H{"commands": commands})
}

func CreateCommand(c *gin.Context) {
	var input models.Command
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&input)
  c.JSON(http.StatusOK, gin.H{"command": input})
}

func ExecuteCommand(c *gin.Context) {
	phrase := c.Param("phrase")
	var command models.Command
	if err := models.DB.Where("phrase = ?", phrase).First(&command).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Command not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"action": command.Action})
}
