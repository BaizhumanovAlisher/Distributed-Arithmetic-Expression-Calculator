package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8099"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	Path                            string `yaml:"path" env-required:"true"`
	DurationInSecondAddition        int    `yaml:"addition" env-default:"0"`
	DurationInSecondSubtraction     int    `yaml:"subtraction" env-default:"0"`
	DurationInSecondAMultiplication int    `yaml:"multiplication" env-default:"0"`
	DurationInSecondDivision        int    `yaml:"division" env-default:"0"`
}

func MustLoad() *Config {
	configPath := "config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
