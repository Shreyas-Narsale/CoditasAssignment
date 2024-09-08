package config

import (
	"os"
)

type Log struct {
	LogMode string
}

var logConfig = &Log{}

func GetLogConfig() *Log {
	return logConfig
}

func LoadLogConfig() {
	logConfig.LogMode = os.Getenv("LOG_MODE")
}
