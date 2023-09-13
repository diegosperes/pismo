package util

import (
	"log"
)

func SetupApp() {
	if settingsErr := LoadSettings(); settingsErr != nil {
		log.Fatal(settingsErr)
	}

	settings := GetSettings()
	logLevelName := settings.Server.LogLevelName

	if logErr := ConfigureDefaultLogger(logLevelName); logErr != nil {
		log.Fatal(logErr)
	}

	if dbErr := SetupDatabaseConnection(settings.Database, logLevelName); dbErr != nil {
		log.Fatal(dbErr)
	}
}
