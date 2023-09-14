package util

import (
	"log"

	"gorm.io/gorm"
)

type AppDependencies struct {
	Settings AppSettings
	Database gorm.DB
}

func SetupApp() *AppDependencies {
	settings, settingsErr := LoadSettings()

	if settingsErr != nil {
		log.Fatal(settingsErr)
	}

	logLevelName := settings.Server.LogLevelName

	if logErr := ConfigureDefaultLogger(logLevelName); logErr != nil {
		log.Fatal(logErr)
	}

	database, dbErr := SetupDatabaseConnection(settings.Database, logLevelName)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	return &AppDependencies{
		Settings: *settings,
		Database: *database,
	}
}
