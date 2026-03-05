package vectorstore

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate/entities/models"
)

const CollectionName = "BotKnowledge"

func (vs *VectorStore) ensureSchema(ctx context.Context) error {
	exists, err := vs.client.Schema().ClassExistenceChecker().
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
		Vectorizer:  "none",
		Description: "Bot knowledge entries for projects",
		Properties: []*models.Property{
			{
				Name:     "projectId",
				DataType: []string{"text"},
			},
			{
				Name:     "knowledgeIndex",
				DataType: []string{"int"},
			},
			{
				Name:     "content",
				DataType: []string{"text"},
			},
		},
	}

	err = vs.client.Schema().ClassCreator().
		WithClass(classDef).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("create schema: %w", err)
	}
	return nil
}
