package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func init() {
	rootCmd.AddCommand(genAPI)
}

var rootCmd = &cobra.Command{
	Use:   "version",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		zap.S().Error("Error when execute command, detail: ", err)
		os.Exit(0)
	}
}
