package vectorstore

import (
	"errors"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

var weaviateClient *weaviate.Client

func InitWeaviate(cfg config.WeaviateConfig) error {
	if cfg.Host == "" {
		return errors.New("WEAVIATE_HOST is not set")
	}
	scheme := cfg.Scheme
	if scheme == "" {
		scheme = "http"
	}
	client, err := weaviate.NewClient(weaviate.Config{
		Host:   cfg.Host,
		Scheme: scheme,
	})
	if err != nil {
		return err
	}
	weaviateClient = client
	return nil
}

func GetWeaviateClient() (*weaviate.Client, error) {
	if weaviateClient == nil {
		return nil, errors.New("Weaviate client not initialized")
	}
	return weaviateClient, nil
}
