package qdrant

import (
	"context"
	"fmt"

	qdrant "github.com/qdrant/go-client/qdrant"
)

type SearchResult struct {
	Filename string
	Chunk    int
	Text     string
	Score    float32
}

func SearchSimilarChunks(
	client *qdrant.Client,
	queryVector []float32,
	limit uint64,
) ([]SearchResult, error) {

	ctx := context.Background()

	// Request payload in the response
	withPayload := qdrant.NewWithPayload(true)

	results, err := client.Query(ctx, &qdrant.QueryPoints{
		CollectionName: CollectionName,
		Query:          qdrant.NewQueryDense(queryVector),
		Limit:          &limit,
		WithPayload:    withPayload,
	})

	if err != nil {
		return nil, err
	}

	var searchResults []SearchResult

	for _, point := range results {

		// Debug output
		fmt.Printf("Score: %f\n", point.Score)
		fmt.Printf("Payload: %+v\n", point.Payload)

		result := SearchResult{
			Score: point.Score,
		}

		if value, ok := point.Payload["filename"]; ok {
			result.Filename = value.GetStringValue()
		}

		if value, ok := point.Payload["chunk"]; ok {
			result.Chunk = int(value.GetIntegerValue())
		}

		if value, ok := point.Payload["text"]; ok {
			result.Text = value.GetStringValue()
		}

		fmt.Printf("Search Result: %+v\n", result)

		searchResults = append(searchResults, result)
	}

	return searchResults, nil
}
