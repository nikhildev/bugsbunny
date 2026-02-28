package commands

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/routes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type serverConfig struct {
	Host string
	Port string
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Start the BugsBunny API server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v := viper.New()
		v.AutomaticEnv()
		v.SetEnvPrefix("HTTP_SERVER")
		v.SetConfigFile(".env")
		v.SetConfigType("env")
		if err := v.ReadInConfig(); err != nil {
			return err
		}
		serverConfig := serverConfig{
			Host: v.GetString("HTTP_SERVER_HOST"),
			Port: v.GetString("HTTP_SERVER_PORT"),
		}

		cfg := clients.GetDbConfig()
		if _, err := clients.InitDB(cfg); err != nil {
			return err
		}

		defer func() {
			slog.Info("Closing database connection")
			clients.CloseDbClient()
		}()

		addr := serverConfig.Host + ":" + serverConfig.Port
		slog.Info("Starting server", "addr", addr)
		mux := routes.SetupRoutes()
		handler := common.Chain(mux, common.RecoveryMiddleware, common.LoggingMiddleware, common.CORSMiddleware)
		return http.ListenAndServe(addr, handler)
	},
}
