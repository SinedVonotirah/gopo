package main

import (
	"github.com/SinedVonotirah/gopo/core/cmd"
	"github.com/SinedVonotirah/gopo/shared/logging"

	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Info("Error at gopo startup")
		os.Exit(1)
	}

}
