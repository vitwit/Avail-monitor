package config

import (
	"log"
	"os/user"
	"path"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
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
		URLEndpoint string `mapstructure:"url_endpoint"`
	}

	Config struct {
		Prometheus Prometheus `mapstructure:"prometheus"`
		Scraper    Scraper    `mapstructure:"scraper"`
		Endpoints  Endpoints  `mapstructure:"url_endpoint"`
	}

	Response struct {
		ClientVersion string `json:"clientVersion"`
	}
)

func ReadFromFile() (*Config, error) {
	usr, err := user.Current()
	if err != nil {
		log.Printf("Error while reading current user : %v", err)
	}

	configPath := path.Join(usr.Name, `./config.toml`)
	log.Printf("config path of root: %s", &configPath)

	v := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
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

	return &cfg, nil

}

func (c *Config) Validate(e ...string) error {
	v := validator.New()
	if len(e) == 0 {
		return v.Struct(c)
	}
	return v.StructExcept(c, e...)
}
