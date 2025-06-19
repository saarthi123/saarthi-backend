package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/utils"
)

func RefreshTokenHandler(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token required"})
		return
	}

	claims, err := utils.VerifyRefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate new access token
	newAccessToken, err := utils.CreateJWTToken(claims.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}



