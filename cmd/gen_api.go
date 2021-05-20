package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var genAPI = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Info("Start command api")

		zap.S().Info("Stop command api")
	},
}

