package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/chunk"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/config"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/embeddings"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/pdf"
	qdrantdb "github.com/ujjwal0563/go-rag-pdf-chat/internal/qdrant"
)

type EmbeddingRequest struct {
	Filename string `json:"filename"`
}

func GenerateEmbeddings(c *gin.Context) {

	var req EmbeddingRequest

	// Parse request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Load config
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

	// Connect to Qdrant
	client, err := qdrantdb.NewClient(cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer client.Close()

	// Create collection if it doesn't exist
	err = qdrantdb.CreateCollection(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Read PDF
	filePath := filepath.Join(cfg.UploadPath, req.Filename)

	text, err := pdf.ReadPDF(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Split text into chunks
	chunks := chunk.SplitText(text, 500)

	var embeddingDimension int
	var totalVectors int

	for index, ch := range chunks {

		// Generate embedding
		vector, err := service.GenerateEmbedding(ch)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		embeddingDimension = len(vector)

		// Store vector in Qdrant
		err = qdrantdb.InsertVector(
			client,
			req.Filename,
			index,
			ch,
			vector,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		totalVectors++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":            "Embeddings generated and stored successfully",
		"embeddingDimension": embeddingDimension,
		"totalChunks":        len(chunks),
		"totalVectors":       totalVectors,
	})
}
