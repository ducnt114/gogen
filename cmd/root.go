package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var dbType string
var dbHost string
var dbPort int32
var dbUser string
var dbPass string
var dbName string
var dbTable string

func init() {
	genAPI.Flags().StringVarP(&dbType, "type", "g", "mysql", "db type: mysql, postgres")
	genAPI.Flags().StringVarP(&dbHost, "host", "o", "localhost", "db host")
	genAPI.Flags().Int32VarP(&dbPort, "port", "p", 3306, "db port")
	genAPI.Flags().StringVarP(&dbUser, "user", "u", "root", "db user")
	genAPI.Flags().StringVarP(&dbPass, "pass", "w", "secret", "db pass")
	genAPI.Flags().StringVarP(&dbName, "db", "d", "db_name", "db name")
	genAPI.Flags().StringVarP(&dbTable, "table", "t", "table_name", "db name")

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
