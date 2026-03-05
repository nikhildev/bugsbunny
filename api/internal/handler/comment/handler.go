package comment

import "github.com/nikhildev/bugsbunny/api/internal/repository"

type Handler struct {
	Repo repository.CommentRepo
}
