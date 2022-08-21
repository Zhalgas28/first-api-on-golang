package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"man_utd/pkg/logging"
	"sync"
)

type Config struct {
	Port string `yaml:"port"`
	DB   struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		DBName   string `yaml:"db_name"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"db"`
}

var once sync.Once
var instance *Config

func GetConfig() *Config {
	logger := logging.GetLogger()
	once.Do(func() {
		instance = &Config{}
		logger.Info("creating config")
		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			info, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(info)
			logger.Fatal(err)
		}
	})
	return instance
}
