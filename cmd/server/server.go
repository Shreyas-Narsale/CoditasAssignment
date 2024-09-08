package server

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"SOCIAL_MEDIA_APP/pkg/logger"
	"SOCIAL_MEDIA_APP/pkg/middleware"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Serve() {
	logs := logger.GetLogger()

	fiberConfig := config.GetFiberConfig()
	app := fiber.New(fiberConfig)
	middleware.FiberMiddleware(app)
	LoadRoutes(app)

	go func() {
		AppConfig := config.GetAppConfig()
		ServerAddr := fmt.Sprintf("%s:%d", AppConfig.HostIp, AppConfig.Port)
		if err := app.Listen(ServerAddr); err != nil {
			logs.Fatal().Err(err).Msg("Error starting server")
		}
	}()

	// Create a channel to listen for OS signals
	shutdownDone := make(chan struct{})

	// Create a channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start a separate goroutine for shutdown handling
	go func() {
		sig := <-stop
		logs.Warn().Msgf("Received signal:%v", sig)

		// Create a context with a timeout for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Gracefully stop the server
		if err := app.ShutdownWithContext(ctx); err != nil {
			logs.Err(err).Msg("Error during shutdown")
		}

		logs.Warn().Msg("Server gracefully stopped")
		close(shutdownDone)
	}()

	// Wait for shutdown signal to complete
	<-shutdownDone
	logs.Warn().Msg("Exiting program")
}
