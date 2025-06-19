package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type AIRequest struct {
    Query string `json:"query"`
}

type AIResponse struct {
    Answer string `json:"answer"`
}

func AskAI(c *gin.Context) {
    var req AIRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    answer := "This is a mock AI answer for: " + req.Query

   
c.JSON(http.StatusOK, gin.H{
    "answer": answer,
})}
