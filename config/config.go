package config

import (
	"github.com/StepanchukYI/user-service/server/http"
	"github.com/StepanchukYI/user-service/server/socket"
	"github.com/StepanchukYI/user-service/server/statuscheck"
)

type Config struct {
	LogLevel string             `mapstructure:"LOG_LEVEL" default:"DEBUG"`
	HTTP     http.Config        `mapstructure:"HTTP"`
	WSS      websocket.Config   `mapstructure:"WSS"`
	STATUS   statuscheck.Config `mapstructure:"STATUS"`
}
