package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/accesscontrol"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/donovanhubbard/wizard-duel/app"
	"github.com/muesli/termenv"
)

const (
	host = "0.0.0.0"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Infof("Starting program.")

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "23234"
	}
	log.Info(fmt.Sprintf("Using port: %s", portStr))
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Error(fmt.Sprintf("could not convert port to int: %v", err))
		os.Exit(1)
	}

	keyPath := os.Getenv("PRIVATE_KEY_PATH")
	if keyPath == "" {
		keyPath = ".ssh/term_info_ed25519"
	}
	log.Info(fmt.Sprintf("Using key path: %s", keyPath))

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		log.Warn(fmt.Sprintf("Key path does not exist: %s. Generating a new key pair", keyPath))
	}

	a := app.NewApp()

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(keyPath),
		wish.WithMiddleware(
			bm.MiddlewareWithProgramHandler(a.ProgramHandler, termenv.ANSI256),
			logging.Middleware(),
			activeterm.Middleware(),
			accesscontrol.Middleware(),
		),
	)

	if err != nil {
		log.Error("could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("could not start server", "error", err)
			done <- nil
		}
	}()

	a.Start()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("could not stop server", "error", err)
	}
}
