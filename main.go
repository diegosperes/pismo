package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/diegosperes/pismo/app/handle"
	"github.com/diegosperes/pismo/app/util"
)

func main() {
	deps := util.SetupApp()
	port := deps.Settings.Server.Port

	slog.Info("Starting application", "port", port)
	serverErr := http.ListenAndServe(port, handle.GetConfiguredRouter(deps))

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
