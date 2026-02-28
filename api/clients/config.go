package clients

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig holds all application configuration.
type AppConfig struct {
	DB     DbConfig
	Server ServerConfig
	OpenAI OpenAIConfig
}

// ServerConfig contains HTTP server settings.
type ServerConfig struct {
	Host string
	Port string
}

// OpenAIConfig contains OpenAI API settings.
type OpenAIConfig struct {
	APIKey string
}

// LoadConfig reads all configuration from environment variables and the .env
// file, returning a fully populated AppConfig or an error.
func LoadConfig() (AppConfig, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	if err := v.ReadInConfig(); err != nil {
		return AppConfig{}, fmt.Errorf("read config: %w", err)
	}
	return AppConfig{
		DB: DbConfig{
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetString("DB_PORT"),
			User:     v.GetString("DB_USER"),
			Password: v.GetString("DB_PASSWORD"),
			Name:     v.GetString("DB_NAME"),
			SSLMode:  v.GetString("DB_SSL_MODE"),
		},
		Server: ServerConfig{
			Host: v.GetString("HTTP_SERVER_HOST"),
			Port: v.GetString("HTTP_SERVER_PORT"),
		},
		OpenAI: OpenAIConfig{
			APIKey: v.GetString("OPENAI_API_KEY"),
		},
	}, nil
}
