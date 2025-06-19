package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

// LoginHandler handles login requests
func LoginHandler(c *gin.Context) {
	var loginPayload struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	// Parse and validate the incoming JSON payload
	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"toast": models.ToastMessage{
				Message: "Invalid request payload",
				Type:    "error",
			},
		})
		return
	}

	// Dummy login validation (replace with actual DB logic)
	if loginPayload.Phone == "9999999999" && loginPayload.Password == "123456" {
		c.JSON(http.StatusOK, gin.H{
			"token": "some-jwt-token", // Replace with real JWT
			"toast": models.ToastMessage{
				Message: "Login successful!",
				Type:    "success",
			},
		})
		return
	}

	// Invalid credentials response
	c.JSON(http.StatusUnauthorized, gin.H{
		"toast": models.ToastMessage{
			Message: "Invalid credentials",
			Type:    "error",
		},
	})
}
