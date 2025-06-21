package main

import (
	"CODITAS_TASK/cmd/server"
	"CODITAS_TASK/pkg/config"
	"CODITAS_TASK/pkg/logger"
	"CODITAS_TASK/platform/database"
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
