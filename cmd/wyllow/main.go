package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wyllow-xyz/wyllow/internal/server"
	"github.com/wyllow-xyz/wyllow/logger"
)

func main() {
	srv := server.New()

	// Create a channel to receive OS signals
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a new goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP Server Failed To Start: %+v\n", err)
		}
	}()
	logger.Info("HTTP Server Listening On %s...", srv.Addr)

	// Wait for the exit signal
	<-exit
	logger.Info("HTTP Server Shutting Down...")

	// Create a context with a timeout for the server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown Failed: %+v\n", err)
	}
	logger.Info("Server Exited Properly\n")
}
