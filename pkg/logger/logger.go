package logger

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Define logging modes and their corresponding zerolog levels
var logModeMap = map[string]zerolog.Level{
	"ALL":      zerolog.DebugLevel,
	"INFO":     zerolog.InfoLevel,
	"WARN":     zerolog.WarnLevel,
	"CRITICAL": zerolog.ErrorLevel,
}

// SetLogLevel sets the global log level based on the selected mode
func SetLogLevel() {

	logConf := config.GetLogConfig()
	mode := logConf.LogMode

	// Default to Debug level if mode is not found
	logLevel := zerolog.DebugLevel

	if level, exists := logModeMap[strings.ToUpper(mode)]; exists {
		logLevel = level
	}
	zerolog.SetGlobalLevel(logLevel)
}

// NewLogger creates and returns a configured zerolog logger
func GetLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()
}
