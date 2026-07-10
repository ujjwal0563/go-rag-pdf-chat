package qdrant

import (
	"context"
	"fmt"

	qdrant "github.com/qdrant/go-client/qdrant"
)

func InsertVector(
	client *qdrant.Client,
	filename string,
	chunkIndex int,
	chunkText string,
	vector []float32,
) error {

	ctx := context.Background()

	// Create payload values
	filenameValue, err := qdrant.NewValue(filename)
	if err != nil {
		return err
	}

	chunkValue, err := qdrant.NewValue(chunkIndex)
	if err != nil {
		return err
	}

	textValue, err := qdrant.NewValue(chunkText)
	if err != nil {
		return err
	}

	// Create point
	point := &qdrant.PointStruct{
		Id:      qdrant.NewIDNum(uint64(chunkIndex + 1)),
		Vectors: qdrant.NewVectors(vector...),
		Payload: map[string]*qdrant.Value{
			"filename": filenameValue,
			"chunk":    chunkValue,
			"text":     textValue,
		},
	}

	wait := true

	// Insert point into Qdrant
	_, err = client.Upsert(ctx, &qdrant.UpsertPoints{
		CollectionName: CollectionName,
		Wait:           &wait,
		Points: []*qdrant.PointStruct{
			point,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to insert vector: %w", err)
	}

	fmt.Printf("Inserted chunk %d successfully\n", chunkIndex)

	return nil
}
