package main

import (
	"log"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/config"
	"github.com/ujjwal0563/go-rag-pdf-chat/internal/router"
)

func main() {

	cfg := config.Load()

	r := router.SetupRouter()

	log.Printf("Server running on port %s", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
