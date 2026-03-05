package project

import (
	"github.com/nikhildev/bugsbunny/api/internal/repository"
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
)

type Handler struct {
	Repo        repository.ProjectRepo
	VectorStore *vectorstore.VectorStore
}
