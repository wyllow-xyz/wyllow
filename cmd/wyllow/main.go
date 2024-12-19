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

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP Server Failed To Start: %+v\n", err)
		}
	}()
	logger.Info("HTTP Server Listening On %s...", srv.Addr)

	<-exit
	logger.Info("HTTP Server Shutting Down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown Failed: %+v\n", err)
	}
	logger.Info("Server Exited Properly\n")
}
