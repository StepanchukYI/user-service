package config

import (
	"stepanchuk/default-template/config/reader"
)

func Read() (Config, error) {
	var cfg Config
	err := reader.Read(&cfg)

	return cfg, err
}
