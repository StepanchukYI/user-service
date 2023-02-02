package statuscheck

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port    int `mapstructure:"PORT"  default:"8081"`
	Timeout int `mapstructure:"TIMEOUT"  default:"60"`
}

type Server struct {
	http     *http.Server
	listener net.Listener
}

func New(cfg Config, statuses ...func() error) *Server {
	return &Server{
		http: &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.Port),
			Handler:           buildHandler(statuses),
			ReadHeaderTimeout: time.Second * time.Duration(cfg.Timeout),
		},
	}
}

func buildHandler(statuses []func() error) http.Handler {
	handler := http.NewServeMux()
	var checks = func(w http.ResponseWriter, _ *http.Request) { serveCheck(w, statuses) }
	handler.HandleFunc("/", checks)

	return handler
}

func serveCheck(writer http.ResponseWriter, checks []func() error) {
	writtenHeader := false
	for _, check := range checks {
		if err := check(); err != nil {
			if !writtenHeader {
				writer.WriteHeader(http.StatusInternalServerError)
				writtenHeader = true
			}
			_, _ = writer.Write([]byte(err.Error()))
			_, _ = writer.Write([]byte("\n\n"))
		}
	}

	if !writtenHeader {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func (s *Server) Serve() error {
	netListen, err := net.Listen("tcp", s.http.Addr)
	if err != nil {
		return fmt.Errorf("faile to create NewWssServer network con: %w", err)
	}
	log.Info(log.WithField("NewHttpServer server at", s.http.Addr))
	s.listener = netListen

	err = s.http.Serve(s.listener)
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.http.Shutdown(ctx)

	return err
}

type Response struct {
	Error string `json:"error"`
}
