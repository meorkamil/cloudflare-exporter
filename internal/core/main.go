package core

import (
	"cloudflare-status/internal/metrics"
	"cloudflare-status/internal/models"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

type Core struct {
	core    *models.Config
	ExpAddr string
	ExpUri  string
}

func NewCloudFlareExp(c string, v string) (*Core, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(c)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config: %s", err)
	}

	var config models.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("Error unmarshalling config: %s", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	config.Version = v

	return &Core{
		core:    &config,
		ExpAddr: config.Exporter.Listen,
		ExpUri:  config.Exporter.Uri,
	}, nil
}

func (c *Core) Run() {
	// Create a channel to listen for shutdown signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a context that can be canceled for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Init Metrics
	m, err := metrics.NewMetrics(c.core.Exporter.Timeout, c.core.Exporter.Interval, c.core.Exporter.API)
	if err != nil {
		log.Fatal(err)
	}

	// Init routers
	mux := http.NewServeMux()
	mux.Handle(c.ExpUri, promhttp.Handler())

	// Init server
	srv := &http.Server{
		Addr:    c.ExpAddr,
		Handler: mux,
	}

	// Start a server in a goroutine
	go func() {
		slog.Info(fmt.Sprintf("Starting Cloudflare Exporter - %s", c.core.Version))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Start record metrics
	go m.RecordMetrics()

	// Wait for shutdown signal or context timeout
	select {
	case <-signalChan:
		slog.Info("Received shutdown signal")
	case <-ctx.Done():
		slog.Info("Shutdown timeout reached")
	}

	// Attemp for graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed:", err)
	}
	slog.Info("Server gracefully stopped")
}
