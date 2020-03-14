package config

import (
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	EventAPI EventAPI `toml:"event_api"`
	DB       Database `toml:"database"`
}

type EventAPI struct {
	Enabled bool   `toml:"enabled"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
}

type Database struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

func (api *EventAPI) Address() string {
	return fmt.Sprintf("nats://%s:%d", api.Host, api.Port)
}

func New(r io.Reader) (*Settings, error) {
	var conf Settings

	if _, err := toml.DecodeReader(r, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
