package config

import (
	"os"
	"strconv"
)

type App struct {
	Port   int
	HostIp string
}

var app = &App{}

func GetAppConfig() *App {
	return app
}

func LoadApplicationConfig() {
	app.HostIp = os.Getenv("HOSTIP")
	app.Port, _ = strconv.Atoi(os.Getenv("PORT"))
}
