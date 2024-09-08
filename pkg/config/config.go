package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

func SetEnvironmentVariable() {
	log.Info().Msg("Jay Ganesh ..")
	filepath := "pkg/config/env-config/environment.json"

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening config file")
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting file info")
	}

	buffer := make([]byte, fileInfo.Size())

	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading file")
	}

	var config map[string]string
	err = json.Unmarshal(buffer, &config)
	if err != nil {
		log.Fatal().Err(err).Msg("Error parsing config file")
	}

	for key, value := range config {
		os.Setenv(key, value)
	}

	log.Info().Msg("Environment are variables set successfully")

}

func LoadConfig() {
	LoadApplicationConfig()
	LoadLogConfig()
	LoadDbConfig()
	LoadGlobalConfig()
}
