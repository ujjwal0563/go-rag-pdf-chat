package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/chunk"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/config"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/embeddings"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/pdf"
)

type EmbeddingRequest struct {
	Filename string `json:"filename"`
}

func GenerateEmbeddings(c *gin.Context) {

	var req EmbeddingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	cfg := config.Load()
	if cfg.GeminiAPIKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gemini API key is missing",
		})
		return
	}
	service, err := embeddings.NewGeminiEmbeddingService(cfg.GeminiAPIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	filePath := filepath.Join(cfg.UploadPath, req.Filename)

	text, err := pdf.ReadPDF(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	chunks := chunk.SplitText(text, 500)

	var vectors int

	for _, ch := range chunks {

		_, err := service.GenerateEmbedding(ch)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		vectors++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Embeddings generated successfully",
		"totalChunks":  len(chunks),
		"totalVectors": vectors,
	})
}
