package helper

import (
	"context"
	"fmt"
	"miniproject/configs"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateDescription(foodName string) (string, error) {
	apiKey := configs.Cfg.GeminiAPIKey
	ctx := context.Background()

	prompt := fmt.Sprintf("Buat deskripsi menarik untuk makanan sisa bernama %s yang akan dijual. Jelaskan kualitas, rasa, dan manfaat dari makanan tersebut.", foodName)

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(0.5)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(100)
	model.ResponseMIMEType = "text/plain"

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	var description string
	for _, part := range resp.Candidates[0].Content.Parts {
		description += fmt.Sprint(part)
	}

	return description, nil
}
