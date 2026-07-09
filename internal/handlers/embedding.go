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

	// Read JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Load configuration
	cfg := config.Load()

	if cfg.GeminiAPIKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gemini API key is missing",
		})
		return
	}

	// Create Gemini embedding service
	service, err := embeddings.NewGeminiEmbeddingService(cfg.GeminiAPIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// PDF path
	filePath := filepath.Join(cfg.UploadPath, req.Filename)

	// Read PDF
	text, err := pdf.ReadPDF(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Split into chunks
	chunks := chunk.SplitText(text, 500)

	var vectors int
	var embeddingDimension int

	// Generate embeddings
	for _, ch := range chunks {

		vector, err := service.GenerateEmbedding(ch)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Save embedding dimension
		embeddingDimension = len(vector)

		vectors++
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"message":            "Embeddings generated successfully",
		"embeddingDimension": embeddingDimension,
		"totalChunks":        len(chunks),
		"totalVectors":       vectors,
	})
}
