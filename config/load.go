package config

import (
	"github.com/StepanchukYI/user-service/config/reader"
)

func Read() (Config, error) {
	var cfg Config
	err := reader.Read(&cfg)

	return cfg, err
}
