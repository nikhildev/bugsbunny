package embedding

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/config"
)

type Client struct {
	baseURL string
	model   string
}

func NewClient(cfg config.EmbeddingConfig) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("EMBEDDING_BASE_URL is not set")
	}
	if cfg.Model == "" {
		return nil, fmt.Errorf("EMBEDDING_MODEL is not set")
	}
	return &Client{baseURL: cfg.BaseURL, model: cfg.Model}, nil
}

func (c *Client) GetEmbedding(ctx context.Context, text string) ([]float32, error) {
	body, err := json.Marshal(map[string]any{
		"model": c.model,
		"input": text,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal embedding request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/embeddings", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create embedding request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call embedding API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("embedding API returned status %d", resp.StatusCode)
	}

	var result struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode embedding response: %w", err)
	}
	if len(result.Data) == 0 {
		return nil, fmt.Errorf("embedding API returned no data")
	}
	return result.Data[0].Embedding, nil
}
