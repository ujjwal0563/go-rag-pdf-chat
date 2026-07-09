package chunk

import "strings"

// SplitText splits a large text into smaller chunks.
func SplitText(text string, chunkSize int) []string {

	// Remove extra spaces
	text = strings.TrimSpace(text)

	// Split text into words
	words := strings.Fields(text)

	var chunks []string

	var currentChunk []string

	currentLength := 0

	for _, word := range words {

		if currentLength+len(word)+1 > chunkSize {

			chunks = append(chunks, strings.Join(currentChunk, " "))

			currentChunk = []string{}
			currentLength = 0
		}

		currentChunk = append(currentChunk, word)

		currentLength += len(word) + 1
	}

	// Add remaining words
	if len(currentChunk) > 0 {
		chunks = append(chunks, strings.Join(currentChunk, " "))
	}

	return chunks
}
