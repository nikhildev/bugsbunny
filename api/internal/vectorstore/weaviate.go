package vectorstore

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

type VectorStore struct {
	client           *weaviate.Client
	embeddingBaseURL string
	embeddingModel   string
}

func NewVectorStore(cfg config.WeaviateConfig, embCfg config.EmbeddingConfig) (*VectorStore, error) {
	if cfg.Host == "" {
		return nil, errors.New("WEAVIATE_HOST is not set")
	}
	if embCfg.BaseURL == "" {
		return nil, errors.New("EMBEDDING_BASE_URL is not set")
	}
	if embCfg.Model == "" {
		return nil, errors.New("EMBEDDING_MODEL is not set")
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
		return nil, fmt.Errorf("create weaviate client: %w", err)
	}

	vs := &VectorStore{
		client:           client,
		embeddingBaseURL: embCfg.BaseURL,
		embeddingModel:   embCfg.Model,
	}

	if err := vs.ensureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure weaviate schema: %w", err)
	}

	return vs, nil
}
