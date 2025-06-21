package server

import (
	"CODITAS_TASK/pkg/config"
	"CODITAS_TASK/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Serve() {
	logs := logger.GetLogger()

	router := NewRouter()
	AppConfig := config.GetAppConfig()
	ServerAddr := fmt.Sprintf("%s:%d", AppConfig.HostIp, AppConfig.Port)

	server := &http.Server{
		Addr:    ServerAddr,
		Handler: router,
	}

	go func() {
		logs.Info().Msgf("Starting server on %s", ServerAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Fatal().Err(err).Msg("Error starting server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logs.Warn().Msg("Shutting down server...")

}
