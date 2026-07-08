package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/services"
)

func UploadPDF(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "PDF file is required",
		})
		return
	}

	filename, err := services.SavePDF(file, "uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "PDF uploaded successfully",
		"filename": filename,
	})
}
