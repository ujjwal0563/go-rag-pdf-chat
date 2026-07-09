package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/pdf"
)

type ReadRequest struct {
	Filename string `json:"filename"`
}

func ReadPDF(c *gin.Context) {

	var req ReadRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	filePath := filepath.Join("uploads", req.Filename)

	text, err := pdf.ReadPDF(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "PDF read successfully",
		"text":    text,
	})
}
