package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfg *Config

var once sync.Once

type Config struct {
	Port string `yaml:"port" env-default:":8000"`
}

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
