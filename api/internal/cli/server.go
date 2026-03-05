package cli

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/nikhildev/bugsbunny/api/internal/database"
	"github.com/nikhildev/bugsbunny/api/internal/handler"
	"github.com/nikhildev/bugsbunny/api/internal/middleware"
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Start the BugsBunny API server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		db, err := database.InitDB(cfg.DB)
		if err != nil {
			return err
		}

		defer func() {
			slog.Info("Closing database connection")
			database.CloseDB(db)
		}()

		var vs *vectorstore.VectorStore
		vs, err = vectorstore.NewVectorStore(cfg.Weaviate, cfg.Embedding)
		if err != nil {
			slog.Warn("Vector store not available, vector features disabled", "error", err)
		} else {
			slog.Info("Vector store initialized", "model", cfg.Embedding.Model)
		}

		addr := cfg.Server.Host + ":" + cfg.Server.Port
		slog.Info("Starting server", "addr", addr)
		mux := handler.SetupRoutes(db, vs)
		h := middleware.Chain(mux, middleware.RecoveryMiddleware, middleware.LoggingMiddleware, middleware.CORSMiddleware)
		return http.ListenAndServe(addr, h)
	},
}
