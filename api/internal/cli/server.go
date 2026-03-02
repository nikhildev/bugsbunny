package cli

import (
	"context"
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
			database.CloseDbClient()
		}()

		if err := vectorstore.InitWeaviate(cfg.Weaviate); err != nil {
			slog.Warn("Weaviate client not available, vector search disabled", "error", err)
		} else {
			slog.Info("Weaviate client initialized")
			if err := vectorstore.EnsureSchema(context.Background()); err != nil {
				slog.Warn("Failed to ensure Weaviate schema", "error", err)
			} else {
				slog.Info("Weaviate schema ready")
			}
		}

		addr := cfg.Server.Host + ":" + cfg.Server.Port
		slog.Info("Starting server", "addr", addr)
		mux := handler.SetupRoutes(db)
		h := middleware.Chain(mux, middleware.RecoveryMiddleware, middleware.LoggingMiddleware, middleware.CORSMiddleware)
		return http.ListenAndServe(addr, h)
	},
}
