package qdrant

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

const (
	CollectionName = "pdf_documents"
	VectorSize     = 3072
)

func CreateCollection(client *qdrant.Client) error {

	ctx := context.Background()

	// Check whether the collection already exists
	exists, err := client.CollectionExists(ctx, CollectionName)
	if err != nil {
		return fmt.Errorf("failed to check collection: %w", err)
	}

	if exists {
		fmt.Println("Qdrant collection already exists.")
		return nil
	}

	// Create the collection
	err = client.CreateCollection(ctx, &qdrant.CreateCollection{
		CollectionName: CollectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     VectorSize,
			Distance: qdrant.Distance_Cosine,
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}

	fmt.Println("Qdrant collection created successfully.")

	return nil
}
