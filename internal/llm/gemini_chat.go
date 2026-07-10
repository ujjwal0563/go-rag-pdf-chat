package llm

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiChatService struct {
	client *genai.Client
	model  string
}

func NewGeminiChatService(apiKey string) (*GeminiChatService, error) {

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &GeminiChatService{
		client: client,
		model:  "gemini-2.5-flash",
	}, nil
}

func (g *GeminiChatService) GenerateAnswer(question, contextText string) (string, error) {

	ctx := context.Background()

	prompt := fmt.Sprintf(`
You are a helpful AI assistant.

Answer ONLY using the information provided in the context below.

If the answer cannot be found in the context, reply exactly:

"I could not find the answer in the uploaded document."

-------------------------
Context:
%s

-------------------------
Question:
%s
`, contextText, question)

	resp, err := g.client.Models.GenerateContent(
		ctx,
		g.model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate answer: %w", err)
	}

	if resp == nil {
		return "", fmt.Errorf("received nil response from Gemini")
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response generated")
	}

	answer := resp.Text()

	if answer == "" {
		return "", fmt.Errorf("empty response from Gemini")
	}

	return answer, nil
}
