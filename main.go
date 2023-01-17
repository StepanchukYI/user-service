package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/StepanchukYI/user-service/config"
	healthServer "github.com/StepanchukYI/user-service/server/healthcheck"
	httpServer "github.com/StepanchukYI/user-service/server/http"
	wssServer "github.com/StepanchukYI/user-service/server/socket"
)

func main() {
	fx.New(
		fx.Provide(
			newConfig,
			newContext,
		),
		fx.Invoke(initLogger),

		fx.Provide(
			NewHTTPServer,
			NewWssServer,
			NewHealthServer,
		),

		fx.Invoke(
			func(*healthServer.Server) {},
			func(*httpServer.Server) {},
			func(*wssServer.Server) {},
		),
	).Run()
}

func newConfig() config.Config {
	// read service cfg from os env
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	return cfg
}

func newContext() context.Context {
	ctx := context.Background()

	return ctx
}

func NewWssServer(lc fx.Lifecycle, cfg config.Config) *wssServer.Server {
	srv := wssServer.New(cfg)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			netListen, err := net.Listen("tcp", srv.GetPort())
			if err != nil {
				return fmt.Errorf("faile to create NewWssServer network con: %w", err)
			}
			log.Info(log.WithField("NewWssServer server at", srv.GetPort()))
			go func() {
				err = srv.Serve(netListen)
				if err != nil {
					log.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)

			return err
		},
	})

	return srv
}

func NewHTTPServer(lc fx.Lifecycle, cfg config.Config) *httpServer.Server {
	srv := httpServer.New(cfg)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			netListen, err := net.Listen("tcp", srv.GetPort())
			if err != nil {
				return fmt.Errorf("faile to create NewHTTPServer network con: %w", err)
			}
			log.Info(log.WithField("NewHttpServer server at", srv.GetPort()))
			go func() {
				err = srv.Serve(netListen)
				if err != nil {
					log.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)

			return err
		},
	})

	return srv
}

func NewHealthServer(lc fx.Lifecycle, cfg config.Config, http *httpServer.Server, wss *wssServer.Server) *healthServer.Server {
	healthChecks := []func() error{http.HealthCheck, wss.HealthCheck}

	srv := healthServer.New(cfg, healthChecks...)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			netListen, err := net.Listen("tcp", srv.GetPort())
			if err != nil {
				return fmt.Errorf("faile to create NewHealthServer network con: %w", err)
			}
			log.Info(log.WithField("NewHealthServer server at", srv.GetPort()))
			go func() {
				err = srv.Serve(netListen)
				if err != nil {
					log.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)

			return err
		},
	})

	return srv
}

func initLogger(cfg config.Config) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	switch strings.ToLower(cfg.LogLevel) {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}
