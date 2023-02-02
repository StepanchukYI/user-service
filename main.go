package main

import (
	"context"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/StepanchukYI/user-service/config"
	httpServer "github.com/StepanchukYI/user-service/server/http"
	wssServer "github.com/StepanchukYI/user-service/server/socket"
	statusServer "github.com/StepanchukYI/user-service/server/statuscheck"
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
			func(*statusServer.Server) {},
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
	srv := wssServer.New(cfg.WSS)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := srv.Serve()
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
	srv := httpServer.New(cfg.HTTP)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := srv.Serve()
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

func NewHealthServer(lc fx.Lifecycle, cfg config.Config, http *httpServer.Server, wss *wssServer.Server) *statusServer.Server {
	statusChecks := []func() error{http.Status, wss.Status}

	srv := statusServer.New(cfg.STATUS, statusChecks...)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := srv.Serve()
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
