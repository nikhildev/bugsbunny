package cli

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/nikhildev/bugsbunny/api/internal/database"
	"github.com/nikhildev/bugsbunny/api/internal/handler"
	"github.com/nikhildev/bugsbunny/api/internal/middleware"
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

		addr := cfg.Server.Host + ":" + cfg.Server.Port
		slog.Info("Starting server", "addr", addr)
		mux := handler.SetupRoutes(db)
		h := middleware.Chain(mux, middleware.RecoveryMiddleware, middleware.LoggingMiddleware, middleware.CORSMiddleware)
		return http.ListenAndServe(addr, h)
	},
}
