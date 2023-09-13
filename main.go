package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/diegosperes/pismo/app/handle"
	"github.com/diegosperes/pismo/app/util"
)

func main() {
	util.SetupApp()

	settings := util.GetSettings()
	slog.Info("Starting application", "port", settings.Server.Port)
	serverErr := http.ListenAndServe(settings.Server.Port, handle.GetConfiguredRouter())

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
