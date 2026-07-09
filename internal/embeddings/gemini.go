package embeddings

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiEmbeddingService struct {
	client *genai.Client
	model  string
}

func NewGeminiEmbeddingService(apiKey string) (*GeminiEmbeddingService, error) {

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &GeminiEmbeddingService{
		client: client,
		model:  "gemini-embedding-001",
	}, nil
}

func (g *GeminiEmbeddingService) GenerateEmbedding(text string) ([]float32, error) {

	ctx := context.Background()

	resp, err := g.client.Models.EmbedContent(
		ctx,
		g.model,
		genai.Text(text),
		&genai.EmbedContentConfig{
			TaskType: "RETRIEVAL_DOCUMENT",
		},
	)

	if err != nil {
		return nil, err
	}

	if len(resp.Embeddings) == 0 {
		return nil, fmt.Errorf("no embedding returned")
	}

	return resp.Embeddings[0].Values, nil
}
