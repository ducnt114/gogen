package main

import (
	"github.com/ducnt114/gogen/cmd"
	"go.uber.org/zap"
)

func init() {
	// init logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger.WithOptions(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}

func main() {
	cmd.Execute()
}
