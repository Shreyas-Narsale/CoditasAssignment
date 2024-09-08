package main

import (
	"SOCIAL_MEDIA_APP/cmd/server"
	"SOCIAL_MEDIA_APP/pkg/config"
	"SOCIAL_MEDIA_APP/pkg/logger"
	"SOCIAL_MEDIA_APP/platform/database"
)

func init() {
	config.SetEnvironmentVariable()
	config.LoadConfig()
	logger.SetLogLevel()
}

func main() {
	database.DatabaseConnections()
	server.Serve()
}
