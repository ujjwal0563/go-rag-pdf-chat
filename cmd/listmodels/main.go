package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/config"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.GeminiAPIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	page, err := client.Models.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print only the first page of models.
	for _, model := range page.Items {
		fmt.Printf("%+v\n", model)
	}
}
