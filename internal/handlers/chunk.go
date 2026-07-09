package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/chunk"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/pdf"
)

type ChunkRequest struct {
	Filename string `json:"filename"`
}

func ChunkPDF(c *gin.Context) {

	var req ChunkRequest

	// Read JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Build file path
	filePath := filepath.Join("uploads", req.Filename)

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

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message":     "PDF chunked successfully",
		"totalChunks": len(chunks),
		"chunks":      chunks,
	})
}
