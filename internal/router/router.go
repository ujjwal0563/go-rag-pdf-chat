package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/handlers"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/health", handlers.Health)

	router.POST("/upload", handlers.UploadPDF)

	router.POST("/read", handlers.ReadPDF)

	return router
}
