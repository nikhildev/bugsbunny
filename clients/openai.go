package clients

import (
	"log"

	"github.com/spf13/viper"

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
	// ctx := context.Background()
	client := openai.NewClient(
		option.WithAPIKey(v.GetString("OPENAI_API_KEY")),
	)
	openaiClient = &client
	return nil
}

func GetOpenAIClient() *openai.Client {
	if openaiClient == nil {
		panic("OpenAI client not initialized")
	}
	return openaiClient
}
