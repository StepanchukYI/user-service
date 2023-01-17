package healthcheck

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/StepanchukYI/user-service/config"
)

type Server struct {
	http *http.Server
}

func New(cfg config.Config, healthChecks ...func() error) *Server {
	return &Server{
		http: &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.HealthPort),
			Handler:           buildHandler(healthChecks),
			ReadHeaderTimeout: time.Second * 60,
		},
	}
}

func buildHandler(healthChecks []func() error) http.Handler {
	handler := http.NewServeMux()
	var checks = func(w http.ResponseWriter, _ *http.Request) { serveCheck(w, healthChecks) }
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

func (s *Server) GetPort() string {
	return s.http.Addr
}

func (s *Server) Serve(ln net.Listener) error {
	err := s.http.Serve(ln)

	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.http.Shutdown(ctx)

	return err
}

type Response struct {
	Error string `json:"error"`
}
