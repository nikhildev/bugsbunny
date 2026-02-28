package commands

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/routes"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Start the BugsBunny API server.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := clients.LoadConfig()
		if err != nil {
			slog.Error("Error loading config", "error", err)
			os.Exit(1)
		}

		db, err := clients.InitDB(cfg.DB)
		if err != nil {
			slog.Error("Error initializing database", "error", err)
			os.Exit(1)
		}
		defer clients.CloseDbClient()

		slog.Info("Starting server", "host", cfg.Server.Host, "port", cfg.Server.Port)
		mux := routes.SetupRoutes(db)
		addr := cfg.Server.Host + ":" + cfg.Server.Port
		if err = http.ListenAndServe(addr, mux); err != nil {
			slog.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	},
}
