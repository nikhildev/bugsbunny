package clients

import (
	"log"

	"github.com/spf13/viper"

	"errors"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

var openaiClient *openai.Client = nil

func InitOpenAI() error {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("OPENAI")
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	apiKey := v.GetString("OPENAI_API_KEY")

	if apiKey == "" {
		return errors.New("OPENAI_API_KEY is not set")
	}
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	openaiClient = &client
	return nil
}

func GetOpenAIClient() (*openai.Client, error) {
	if openaiClient == nil {
		return nil, errors.New("OpenAI client not initialized")
	}
	return openaiClient, nil
}
