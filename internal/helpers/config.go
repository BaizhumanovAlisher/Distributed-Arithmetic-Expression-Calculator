package helpers

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env                string `yaml:"env" env-default:"local"`
	HTTPServer         `yaml:"http_server"`
	Storage            `yaml:"storage"`
	Operation          `yaml:"operation"`
	QuickAccessStorage `yaml:"quick_access_storage"`
	Agent              `yaml:"agent"`
	AuthService        `yaml:"auth_service"`
	GRPCConfig         `yaml:"grpc"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8099"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	User     string `yaml:"user"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"post"`
	SSLMode  string `yaml:"sslmode"`
}

type Operation struct {
	DurationInSecondAddition       int `yaml:"addition" env-default:"0"`
	DurationInSecondSubtraction    int `yaml:"subtraction" env-default:"0"`
	DurationInSecondMultiplication int `yaml:"multiplication" env-default:"0"`
	DurationInSecondDivision       int `yaml:"division" env-default:"0"`
	CountOperation                 int `yaml:"count_operation"`
}

type QuickAccessStorage struct {
	Address  string        `yaml:"addr"`
	TTL      time.Duration `yaml:"ttl"`
	Password string        `yaml:""`
	DB       int           `yaml:"db"`
}

type Agent struct {
	CountOperation int `yaml:"count_calculator"`
}

type AuthService struct {
	Address  string        `yaml:"address" env-default:"localhost:8102"`
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port int `yaml:"port" env-default:"8103"`
}

func MustLoad() *Config {
	configPath := "./config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	checkDuration(&cfg)

	return &cfg
}

func checkDuration(cfg *Config) {
	if cfg.DurationInSecondAddition < 0 {
		log.Fatalf("duration of addition operation is lower than 0")
	}

	if cfg.DurationInSecondSubtraction < 0 {
		log.Fatalf("duration of subtraction operation is lower than 0")
	}

	if cfg.DurationInSecondMultiplication < 0 {
		log.Fatalf("duration of multiplication operation is lower than 0")
	}

	if cfg.DurationInSecondDivision < 0 {
		log.Fatalf("duration of division operation is lower than 0")
	}
}
