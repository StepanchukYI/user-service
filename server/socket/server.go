package websocket

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"stepanchuk/default-template/config"
)

var ErrServiceNotReady = errors.New("wss service: not started yet")

type Server struct {
	http   *http.Server
	ready  bool
	runErr error
}

func New(cfg config.Config) *Server {
	httpSrv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.WSSPort),
		ReadHeaderTimeout: time.Second * 60,
	}

	server := &Server{
		http: httpSrv,
	}

	return server
}

func (s *Server) HealthCheck() error {
	if !s.ready {
		return ErrServiceNotReady
	}
	if s.runErr != nil {
		return s.runErr
	}

	return nil
}

func (s *Server) GetPort() string {
	return s.http.Addr
}

func (s *Server) Serve(ln net.Listener) error {
	s.ready = true
	err := s.http.Serve(ln)
	s.runErr = err

	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.http.Shutdown(ctx)
	s.runErr = err

	return err
}
