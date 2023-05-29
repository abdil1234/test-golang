package config

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type AppConfiguration struct {
	Host    string
	Env     string
	Port    int
	Code    string
	Version string
}

type DatabaseConfiguration struct {
	Driver               string
	Name                 string
	User                 string
	Password             string
	Host                 string
	Port                 int
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

type ErrorDesc struct {
	Id string
	En string
}

var configuration *Configuration
var once sync.Once

type Configuration struct {
	App      AppConfiguration
	Database DatabaseConfiguration
	ErrorMap map[string]ErrorDesc
}

func GetConfig() *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../../.")
		viper.AddConfigPath("../../../.")
		viper.AddConfigPath("../../../../.")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file (abdil), %s", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})

	return configuration
}
