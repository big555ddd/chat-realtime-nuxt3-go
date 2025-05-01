package cmd

import (
	"app/app/routes"
	"app/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// HTTP is serve http or https
func HttpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "Run server on HTTP protocol",
		Run: func(cmd *cobra.Command, args []string) {
			port := viper.GetString("PORT")
			if port == "" {
				port = "8080" // Default to port 8080 if PORT is not set
			}

			r := gin.Default()
			routes.Router(r)
			logger.Infof("Starting server on port %s...", port)
			if err := r.Run(":" + port); err != nil {
				logger.Errf("Failed to start server: %v", err)
			}
		},
	}
}
