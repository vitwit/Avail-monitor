package config

import (
	"log"
	"os/user"
)

type (
	Prometheus struct {
		ListenAddress     string `mapstructure:"listen_address"`
		PrometheusAddress string `mapstructure:"prometheus_address"`
	}

	Scraper struct {
		Rate string `mapstructure:"rate"`
	}

	Endpoints struct {
		WSEndpoint string `mapstructure:"websocket_endpoint"`
	}

	Config struct {
		Prometheus Prometheus `mapstructure:"prometheus"`
		Scraper    Scraper    `mapstructure:"scraper"`
		Endpoints  Endpoints  `mapstructure:"websocket_endpoint"`
	}
)

func ReadFromFile() (*Config, error) {
	usr, err := user.Current()
	if err != nil {
		log.Printf("Error %v", err)
	}
}
