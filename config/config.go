package config

import "github.com/spf13/viper"

type Prometheus struct {
	// ListenAddress to export metrics on the given port
	ListenAddress string `mapstructure:"listen_address"`
	// PrometheusAddress to connect to prormetheus where it has running
	PrometheusAddress string `mapstructure:"prometheus_address"`
}

type Endpoints struct {
	URLEndpoint string `mapstructure:"url_endpoint"`
}

type Config struct {
	URLEndpoint string     `mapstructure:"url_endpoint"`
	Prometheus  Prometheus `mapstructure:"prometheus"`
}

func ReadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// import (
// 	"os"

// 	"github.com/BurntSushi/toml"

// 	// "github.com/spf13/viper/internal/encoding/toml"
// 	"gopkg.in/go-playground/validator.v9"
// )

// type (
// 	Prometheus struct {
// 		ListenAddress     string `mapstructure:"listen_address"`
// 		PrometheusAddress string `mapstructure:"prometheus_address"`
// 	}

// 	// Scraper struct {
// 	// 	Rate string `mapstructure:"rate"`
// 	// }

// 	Endpoints struct {
// 		URLEndpoint string `mapstructure:"url_endpoint"`
// 	}

// 	Config struct {
// 		Prometheus Prometheus `mapstructure:"prometheus"`
// 		// Scraper    Scraper    `mapstructure:"scraper"`
// 		Endpoints Endpoints `mapstructure:"url_endpoint"`
// 	}

// 	Response struct {
// 		ClientVersion string `json:"clientVersion"`
// 	}
// )

// // func ReadFromFile() (*Config, error) {
// // 	// version, err := os.ReadFile("config.toml")

// // if err != nil {
// // 	log.Fatal(err)
// // }
// // fmt.Println(string(version))

// // v := viper.New()
// // v.AddConfigPath(".")
// // v.AddConfigPath("./config/")
// // v.SetConfigName("config")
// // if err := v.ReadInConfig(); err != nil {
// // 	log.Fatalf("error while reading config.toml: %v", err)
// // }
// // var cfg Config
// // if err := v.Unmarshal(&cfg); err != nil {
// // 	log.Fatalf("error unmarshaling config.toml to application config: %v", err)
// // }
// // if err := cfg.Validate(); err != nil {
// // 	log.Fatalf("error occurred in config validation: %v", err)
// // }
// // return &cfg, nil

// // v := viper.New()
// // v.SetConfigType("toml")
// // viper.SetConfigName("config")
// // viper.AddConfigPath(currentUser.HomeDir + "avail-monitor")
// // if err := v.ReadInConfig(); err != nil {
// // 	log.Fatalf("error while reading config.toml: %v", err)
// // }

// // var cfg Config
// // if err := v.Unmarshal(&cfg); err != nil {
// // 	log.Fatalf("error unmarshaling config.toml to application config: %v", err)
// // }

// // if err := cfg.Validate(); err != nil {
// // 	log.Fatalf("error occurred in config validation: %v", err)
// // }

// // return &cfg, nil

// // }

// func (c *Config) GetConfig(configFileName string) error {
// 	// Open the TOML configuration file
// 	tomlFile, err := os.Open(configFileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer tomlFile.Close()

// 	// Decode the TOML content into the Config struct
// 	if _, err := NewDecoder(tomlFile).value(c) err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *Config) Validate(e ...string) error {
// 	v := validator.New()
// 	if len(e) == 0 {
// 		return v.Struct(c)
// 	}
// 	return v.StructExcept(c, e...)
// }
