package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig holds all application configuration.
type AppConfig struct {
	DB        DbConfig
	Server    ServerConfig
	OpenAI    OpenAIConfig
	Weaviate  WeaviateConfig
	Embedding EmbeddingConfig
}

// DbConfig contains database connection settings.
type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
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

// WeaviateConfig contains Weaviate vector store settings.
type WeaviateConfig struct {
	Host   string
	Scheme string
}

// EmbeddingConfig contains settings for the embedding model API.
type EmbeddingConfig struct {
	BaseURL string
	Model   string
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
		Weaviate: WeaviateConfig{
			Host:   v.GetString("WEAVIATE_HOST"),
			Scheme: v.GetString("WEAVIATE_SCHEME"),
		},
		Embedding: EmbeddingConfig{
			BaseURL: v.GetString("EMBEDDING_BASE_URL"),
			Model:   v.GetString("EMBEDDING_MODEL"),
		},
	}, nil
}

// GetDbConfig reads database configuration from environment variables and the
// .env file.
func GetDbConfig() (DbConfig, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("DB")
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	if err := v.ReadInConfig(); err != nil {
		return DbConfig{}, fmt.Errorf("failed to read database config: %w", err)
	}
	return DbConfig{
		Host:     v.GetString("DB_HOST"),
		Port:     v.GetString("DB_PORT"),
		User:     v.GetString("DB_USER"),
		Password: v.GetString("DB_PASSWORD"),
		Name:     v.GetString("DB_NAME"),
		SSLMode:  v.GetString("DB_SSL_MODE"),
	}, nil
}
