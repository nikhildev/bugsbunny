package commands

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/routes"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Start the BugsBunny API server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := clients.LoadConfig()
		if err != nil {
			return err
		}

		if _, err := clients.InitDB(cfg.DB); err != nil {
			return err
		}

		defer func() {
			slog.Info("Closing database connection")
			clients.CloseDbClient()
		}()

		addr := cfg.Server.Host + ":" + cfg.Server.Port
		slog.Info("Starting server", "addr", addr)
		mux := routes.SetupRoutes()
		handler := common.Chain(mux, common.RecoveryMiddleware, common.LoggingMiddleware, common.CORSMiddleware)
		return http.ListenAndServe(addr, handler)
	},
}
