package vectorstore

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/weaviate/weaviate-go-client/v5/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v5/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

var embeddingBaseURL string
var embeddingModel string

func InitEmbeddingClient(cfg config.EmbeddingConfig) error {
	if cfg.BaseURL == "" {
		return errors.New("EMBEDDING_BASE_URL is not set")
	}
	if cfg.Model == "" {
		return errors.New("EMBEDDING_MODEL is not set")
	}
	embeddingBaseURL = cfg.BaseURL
	embeddingModel = cfg.Model
	return nil
}

func getEmbedding(ctx context.Context, text string) ([]float32, error) {
	if embeddingBaseURL == "" {
		return nil, errors.New("embedding client not initialized")
	}

	body, err := json.Marshal(map[string]any{
		"model": embeddingModel,
		"input": text,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal embedding request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, embeddingBaseURL+"/embeddings", bytes.NewReader(body))
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

type SearchResult struct {
	ProjectID      string  `json:"project_id"`
	KnowledgeIndex int     `json:"knowledge_index"`
	Content        string  `json:"content"`
	Distance       float64 `json:"distance"`
}

func SyncProjectKnowledge(ctx context.Context, projectID string, knowledge []string) error {
	client, err := GetWeaviateClient()
	if err != nil {
		return err
	}

	// Delete existing entries for this project
	_, err = client.Batch().ObjectsBatchDeleter().
		WithClassName(CollectionName).
		WithWhere(filters.Where().
			WithPath([]string{"projectId"}).
			WithOperator(filters.Equal).
			WithValueString(projectID)).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("delete existing knowledge: %w", err)
	}

	if len(knowledge) == 0 {
		return nil
	}

	batcher := client.Batch().ObjectsBatcher()
	for i, entry := range knowledge {
		vector, err := getEmbedding(ctx, entry)
		if err != nil {
			return fmt.Errorf("embed knowledge entry %d: %w", i, err)
		}
		batcher.WithObjects(&models.Object{
			Class: CollectionName,
			Properties: map[string]any{
				"projectId":      projectID,
				"knowledgeIndex": i,
				"content":        entry,
			},
			Vector: models.C11yVector(vector),
		})
	}

	resp, err := batcher.Do(ctx)
	if err != nil {
		return fmt.Errorf("batch insert knowledge: %w", err)
	}
	for _, r := range resp {
		if r.Result != nil && r.Result.Errors != nil {
			return fmt.Errorf("batch insert error: %v", r.Result.Errors)
		}
	}
	return nil
}

func SearchKnowledge(ctx context.Context, query string, topK int, projectID string) ([]SearchResult, error) {
	client, err := GetWeaviateClient()
	if err != nil {
		return nil, err
	}

	queryVector, err := getEmbedding(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("embed query: %w", err)
	}

	fields := []graphql.Field{
		{Name: "projectId"},
		{Name: "knowledgeIndex"},
		{Name: "content"},
		{Name: "_additional", Fields: []graphql.Field{{Name: "distance"}}},
	}

	builder := client.GraphQL().Get().
		WithClassName(CollectionName).
		WithFields(fields...).
		WithNearVector(client.GraphQL().NearVectorArgBuilder().WithVector(queryVector)).
		WithLimit(topK)

	if projectID != "" {
		builder = builder.WithWhere(filters.Where().
			WithPath([]string{"projectId"}).
			WithOperator(filters.Equal).
			WithValueString(projectID))
	}

	result, err := builder.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("search knowledge: %w", err)
	}

	if result.Errors != nil {
		return nil, fmt.Errorf("graphql errors: %v", result.Errors)
	}

	data, ok := result.Data["Get"].(map[string]any)
	if !ok {
		return nil, nil
	}
	items, ok := data[CollectionName].([]any)
	if !ok {
		return nil, nil
	}

	var results []SearchResult
	for _, item := range items {
		m, ok := item.(map[string]any)
		if !ok {
			continue
		}
		r := SearchResult{
			ProjectID: getString(m, "projectId"),
			Content:   getString(m, "content"),
		}
		if idx, ok := m["knowledgeIndex"].(float64); ok {
			r.KnowledgeIndex = int(idx)
		}
		if additional, ok := m["_additional"].(map[string]any); ok {
			if dist, ok := additional["distance"].(float64); ok {
				r.Distance = dist
			}
		}
		results = append(results, r)
	}
	return results, nil
}

func GetVectorForText(ctx context.Context, text string) ([]float32, error) {
	return getEmbedding(ctx, text)
}

func getString(m map[string]any, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}
