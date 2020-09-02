package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	EventAPI EventAPI `yaml:"event_api"`
	DB       Database `yaml:"database"`
}

type EventAPI struct {
	Enabled bool   `yaml:"enabled"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

func (api *EventAPI) Address() string {
	return fmt.Sprintf("nats://%s:%d", api.Host, api.Port)
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	var config Settings

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
