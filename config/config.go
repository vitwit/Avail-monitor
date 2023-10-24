package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type Prometheus struct {
	// ListenAddress to export metrics on the given port
	ListenAddress string `mapstructure:"listen_address"`
	// PrometheusAddress to connect to prormetheus where it has running
	PrometheusAddress string `mapstructure:"prometheus_address"`
}

type Endpoint struct {
	URLEndpoint string `mapstructure:"url_endpoint"`
}

type Config struct {
	RPC_Endpoint Endpoint   `mapstructure:"url_endpoint"`
	Prometheus   Prometheus `mapstructure:"prometheus"`
}

func ReadConfig() (*Config, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("./config/")
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config.toml: %v", err)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshaling config.toml to application config: %v", err)
	}
	if err := cfg.Validate(); err != nil {
		log.Fatalf("error occurred in config validation: %v", err)
	}
	fmt.Println("config.....", cfg)
	return &cfg, nil
}

func (c *Config) Validate(e ...string) error {
	v := validator.New()
	if len(e) == 0 {
		return v.Struct(c)
	}
	return v.StructExcept(c, e...)
}
