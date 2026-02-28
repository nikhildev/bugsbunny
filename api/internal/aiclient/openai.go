package aiclient

import (
	"errors"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

var openaiClient *openai.Client = nil

func InitOpenAI(cfg config.OpenAIConfig) error {
	if cfg.APIKey == "" {
		return errors.New("OPENAI_API_KEY is not set")
	}
	client := openai.NewClient(option.WithAPIKey(cfg.APIKey))
	openaiClient = &client
	return nil
}

func GetOpenAIClient() (*openai.Client, error) {
	if openaiClient == nil {
		return nil, errors.New("OpenAI client not initialized")
	}
	return openaiClient, nil
}
