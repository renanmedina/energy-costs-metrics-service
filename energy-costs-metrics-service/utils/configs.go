package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	ENVIRONMENT                  string
	SERVICE_NAME                 string
	PORT                         string
	LOG_FORMAT                   string
	OPEN_TELEMETRY_COLLECTOR_URL string
}

var applicationConfigs Configs

func init() {
	loadConfigs()
}

func GetConfigs() Configs {
	return applicationConfigs
}

func loadConfigs() {
	godotenv.Load()

	applicationConfigs = Configs{
		ENVIRONMENT:                  os.Getenv("ENVIRONMENT"),
		SERVICE_NAME:                 os.Getenv("SERVICE_NAME"),
		PORT:                         os.Getenv("PORT"),
		LOG_FORMAT:                   os.Getenv("LOG_FORMAT"),
		OPEN_TELEMETRY_COLLECTOR_URL: os.Getenv("OPEN_TELEMETRY_COLLECTOR_URL"),
	}
}
