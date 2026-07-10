package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/config"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/embeddings"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/llm"
	qdrantdb "github.com/ujjwal0563/go-rag-pdf-chat/internal/qdrant"
)

type AskRequest struct {
	Question string `json:"question"`
}

func AskQuestion(c *gin.Context) {

	var req AskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	cfg := config.Load()

	// Gemini Embedding Service
	embeddingService, err := embeddings.NewGeminiEmbeddingService(cfg.GeminiAPIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Gemini Chat Service
	chatService, err := llm.NewGeminiChatService(cfg.GeminiAPIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Qdrant Client
	client, err := qdrantdb.NewClient(cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer client.Close()

	// Generate embedding for question
	queryVector, err := embeddingService.GenerateEmbedding(req.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Search similar chunks
	results, err := qdrantdb.SearchSimilarChunks(
		client,
		queryVector,
		5,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var contextBuilder strings.Builder

	var sources []gin.H

	for _, r := range results {

		contextBuilder.WriteString(r.Text)
		contextBuilder.WriteString("\n\n")

		sources = append(sources, gin.H{
			"filename": r.Filename,
			"chunk":    r.Chunk,
			"score":    r.Score,
		})
	}

	answer, err := chatService.GenerateAnswer(
		req.Question,
		contextBuilder.String(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"question": req.Question,
		"answer":   answer,
		"sources":  sources,
	})
}
