package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GoTalkie/go-talkie-broker/internal/server"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	webPort                = "80"
	secondsToKeepConnAlive = 10
)

func main() {
	errC, err := run()
	if err != nil {
		log.Fatalf("Could not run: %s", err)
	}

	if err = <-errC; err != nil {
		log.Fatalf("Error while broker was running: %s", err)
	}

}

func run() (<-chan error, error) {
	ctxMsg := "running broker"
	errC := make(chan error, 1)

	// initializing logger
	l, err := zap.NewProduction() // TODO: Configure zap logger
	if err != nil {
		log.Fatalf("Error creating logger: %s", err)
		return nil, errors.Wrapf(err, ctxMsg)
	}

	// initializing server
	app := server.Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	go func() {
		l.Sugar().Infof("Starting broker service on port %s", webPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	// graceful shutdown
	go func() {
		// notify in case user press ctrl+c
		c := make(chan os.Signal, 1)
		signal.Notify(
			c,
			os.Interrupt,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)

		// block goroutine until we receive termination signal
		sig := <-c
		l.Sugar().Infof("Shutdown signal received: ", sig)

		// create context that will cancel in secondsToKeepConnAlive seconds, keeping connection alive for secondsToKeepConnAlive seconds to fulfill their tasks
		ctx, cancel := context.WithTimeout(context.Background(), secondsToKeepConnAlive*time.Second)
		defer func() {
			l.Sync()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			errC <- err
		}
		l.Info("Shutdown completed")
	}()

	return errC, nil
}
