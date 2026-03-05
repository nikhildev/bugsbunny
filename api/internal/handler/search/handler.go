package search

import "github.com/nikhildev/bugsbunny/api/internal/vectorstore"

type Handler struct {
	VectorStore *vectorstore.VectorStore
}
