package vectorstore

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate/entities/models"
)

const CollectionName = "BotKnowledge"

func EnsureSchema(ctx context.Context) error {
	client, err := GetWeaviateClient()
	if err != nil {
		return err
	}

	exists, err := client.Schema().ClassExistenceChecker().
		WithClassName(CollectionName).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("check schema existence: %w", err)
	}
	if exists {
		return nil
	}

	classDef := &models.Class{
		Class:       CollectionName,
		Vectorizer:  "text2vec-transformers",
		Description: "Bot knowledge entries for components",
		ModuleConfig: map[string]any{
			"text2vec-transformers": map[string]any{
				"vectorizeClassName": false,
			},
		},
		Properties: []*models.Property{
			{
				Name:     "componentId",
				DataType: []string{"text"},
				ModuleConfig: map[string]any{
					"text2vec-transformers": map[string]any{
						"skip": true,
					},
				},
			},
			{
				Name:     "knowledgeIndex",
				DataType: []string{"int"},
				ModuleConfig: map[string]any{
					"text2vec-transformers": map[string]any{
						"skip": true,
					},
				},
			},
			{
				Name:     "content",
				DataType: []string{"text"},
			},
		},
	}

	err = client.Schema().ClassCreator().
		WithClass(classDef).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("create schema: %w", err)
	}
	return nil
}
