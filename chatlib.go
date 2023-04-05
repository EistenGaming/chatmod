package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// GPT
// Initilize the GPT client
func gptInit() *openai.Client {
	return openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}

// Ask GPT something and return the reply to the user
func gptComplete(client *openai.Client, msg string) string {

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "[GPT Error]"
	}

	return resp.Choices[0].Message.Content
}
