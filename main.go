package main

import (
	"github.com/sirupsen/logrus"
	"homekit-server/internal/config"
	"homekit-server/internal/homekit"
	"homekit-server/internal/restapi"
	"os"
	"os/signal"
	"syscall"
)

var BuildVersion = "0.0.1.dev"

func main() {
	logrus.Infof("HomeKit Server %s", BuildVersion)
	logrus.SetLevel(logrus.DebugLevel)

	// Load configuration
	cfg, err := config.Init()

	if err != nil {
		logrus.Error(err)
		panic(0)
	}

	// mqtt

	// Create new homekit server
	hk := homekit.New(&cfg.HomeKit)

	// Gracefully shutdown services
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		_ = <-c
		logrus.Info("Gracefully shutting down...")
		hk.Shutdown()
		restapi.Shutdown()
	}()

	// Start services
	go hk.Start(cfg.Entities)
	restapi.Start(&cfg.Server)
}
