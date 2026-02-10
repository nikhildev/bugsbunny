package routes

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/cmd/api/routes/component"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	RegisterHealthRoutes(mux)
	component.RegisterComponentRoutes(mux)
	return mux
}
