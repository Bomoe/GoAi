package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type OpenAiRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Temperature float64 `json:"temperature"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content string `json:"content"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: goai \"your query here\"")
	}

	query := strings.Join(os.Args[1:], " ")
	client := &http.Client{}

	resp, err := sendOpenAiReq("gpt-4o-mini", query, client)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}

func goDotEnvVariable(key string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	err = godotenv.Load(filepath.Join(home, ".goai.env"))
	if err != nil {
		return "", errors.New("Error loading .env file")
	}
	return os.Getenv(key), nil
}

func sendOpenAiReq(model, query string, client *http.Client) (msg string, err error) {
	var allowedModels = []string{"gpt-4o-mini"}
	var isModelAllowed = false
	for _, v := range allowedModels {
		if v == model {
			isModelAllowed = true
			break
		}
	}
	if !isModelAllowed {
		return "", errors.New("Model is not valid, please try again with a valid model.")
	}
	openAiKey, err := goDotEnvVariable("OPENAI_API_KEY")
	if err != nil {
		return "", errors.New(err.Error())
	}
	openAiSettings := OpenAiRequest{
		Model: model,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: query,
			},
		},
		Temperature: 0.7,
	}
	reqBody, err := json.Marshal(openAiSettings)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", openAiKey))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var formattedBody Response
	jsonError := json.Unmarshal([]byte(body), &formattedBody)
	if jsonError != nil {
		return "", jsonError
	}

	if len(formattedBody.Choices) > 0 {
		return formattedBody.Choices[0].Message.Content, nil
	}

	return "", errors.New("No response from OpenAi API found.")
}
