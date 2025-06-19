package models

type AIRequest struct {
    Prompt   string `json:"prompt" binding:"required"`
    Category string `json:"category" binding:"required"`
}

type AIResponse struct {
    Result string `json:"result"`
}

type AiTipResponse struct {
    Tips []string `json:"tips"`
}

// models/ai.go


type AISuggestionRequest struct {
	Context string `json:"context"`
}

