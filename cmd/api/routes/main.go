package routes

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/cmd/api/routes/component"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("POST /components", component.CreateComponentHandler)
	return mux
}
