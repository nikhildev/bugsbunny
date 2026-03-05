package vectorstore

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/nikhildev/bugsbunny/api/internal/embedding"
	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

type VectorStore struct {
	client   *weaviate.Client
	embedder *embedding.Client
}

func NewVectorStore(cfg config.WeaviateConfig, embedder *embedding.Client) (*VectorStore, error) {
	if cfg.Host == "" {
		return nil, errors.New("WEAVIATE_HOST is not set")
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
		client:   client,
		embedder: embedder,
	}

	if err := vs.ensureSchema(context.Background()); err != nil {
		return nil, fmt.Errorf("ensure weaviate schema: %w", err)
	}

	return vs, nil
}
