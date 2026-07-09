package qdrant

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/qdrant/go-client/qdrant"

	"github.com/ujjwal0563/go-rag-pdf-chat/internal/config"
)

func NewClient(cfg *config.Config) (*qdrant.Client, error) {

	if cfg.QdrantURL == "" {
		return nil, fmt.Errorf("QDRANT_URL is empty")
	}

	if cfg.QdrantAPIKey == "" {
		return nil, fmt.Errorf("QDRANT_API_KEY is empty")
	}

	u, err := url.Parse(cfg.QdrantURL)
	if err != nil {
		return nil, fmt.Errorf("invalid Qdrant URL: %w", err)
	}

	port := 6334

	if u.Port() != "" {
		p, err := strconv.Atoi(u.Port())
		if err == nil {
			port = p
		}
	}

	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   u.Hostname(),
		Port:   port,
		APIKey: cfg.QdrantAPIKey,
		UseTLS: u.Scheme == "https",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Qdrant: %w", err)
	}

	return client, nil
}
