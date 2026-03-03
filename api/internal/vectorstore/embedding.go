package vectorstore

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v5/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v5/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

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

	// Insert new entries — Weaviate auto-embeds the content property
	batcher := client.Batch().ObjectsBatcher()
	for i, entry := range knowledge {
		batcher.WithObjects(&models.Object{
			Class: CollectionName,
			Properties: map[string]any{
				"projectId":      projectID,
				"knowledgeIndex": i,
				"content":        entry,
			},
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

	fields := []graphql.Field{
		{Name: "projectId"},
		{Name: "knowledgeIndex"},
		{Name: "content"},
		{Name: "_additional", Fields: []graphql.Field{{Name: "distance"}}},
	}

	builder := client.GraphQL().Get().
		WithClassName(CollectionName).
		WithFields(fields...).
		WithNearText(client.GraphQL().NearTextArgBuilder().WithConcepts([]string{query})).
		WithLimit(topK)

	if projectID != "" {
		builder = builder.WithWhere(filters.Where().
			WithPath([]string{"projectId"}).
			WithOperator(filters.Equal).
			WithValueString(projectID))
	}

	resp, err := builder.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("search knowledge: %w", err)
	}

	if resp.Errors != nil {
		return nil, fmt.Errorf("graphql errors: %v", resp.Errors)
	}

	data, ok := resp.Data["Get"].(map[string]any)
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
	client, err := GetWeaviateClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Batch().ObjectsBatcher().
		WithObjects(&models.Object{
			Class: CollectionName,
			Properties: map[string]any{
				"projectId":      "_simulate_",
				"knowledgeIndex": 0,
				"content":        text,
			},
		}).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("create vectorization object: %w", err)
	}
	if len(resp) == 0 {
		return nil, fmt.Errorf("batch insert returned no response")
	}
	if resp[0].Result != nil && resp[0].Result.Errors != nil {
		return nil, fmt.Errorf("batch insert error: %v", resp[0].Result.Errors)
	}

	id := resp[0].Object.ID.String()

	defer func() {
		_ = client.Data().Deleter().
			WithClassName(CollectionName).
			WithID(id).
			Do(context.Background())
	}()

	objs, err := client.Data().ObjectsGetter().
		WithClassName(CollectionName).
		WithID(id).
		WithVector().
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve vectorized object: %w", err)
	}
	if len(objs) == 0 || objs[0].Vector == nil {
		return nil, fmt.Errorf("no vector returned for object")
	}

	return objs[0].Vector, nil
}

func getString(m map[string]any, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}
