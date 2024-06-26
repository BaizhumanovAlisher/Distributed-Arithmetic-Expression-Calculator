package helpers

import (
	"fmt"
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
	ExpressionSolver   `yaml:"expression_solver"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8099"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	User        string `yaml:"user"`
	DBName      string `yaml:"dbname"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	SSLMode     string `yaml:"sslmode"`
	StoragePath string
}

type Operation struct {
	DurationInSecondAddition       int  `yaml:"addition" env-default:"0"`
	DurationInSecondSubtraction    int  `yaml:"subtraction" env-default:"0"`
	DurationInSecondMultiplication int  `yaml:"multiplication" env-default:"0"`
	DurationInSecondDivision       int  `yaml:"division" env-default:"0"`
	CountOperation                 int  `yaml:"count_operation" env-default:"4"`
	PermissionToSeed               bool `yaml:"permission_to_seed" env-default:"false"`
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
	GrpcPort    int           `yaml:"grpc_port" env-default:"8102"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"1h"`
	Secret      string        `yaml:"secret"`
	CheckSecret bool          `yaml:"check_required_secret" env-default:"false"`
	Cost        int           `yaml:"cost"`
	Path        string        `yaml:"path"`
}

type ExpressionSolver struct {
	GrpcPort int    `yaml:"grpc_port" env-default:"8103"`
	Path     string `yaml:"path"`
}

func MustLoadConfig() *Config {
	configPath := "./config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	cfg.checkDuration()
	cfg.compileStoragePath()
	cfg.checkRequiredSecret()

	return &cfg
}

func (cfg *Config) checkDuration() {
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

func (cfg *Config) compileStoragePath() {
	cfg.Storage.StoragePath = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Storage.User,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.DBName,
		cfg.Storage.SSLMode,
	)
}

func (cfg *Config) checkRequiredSecret() {
	if cfg.AuthService.CheckSecret {
		if len(cfg.AuthService.Secret) == 0 {
			log.Fatalf("secret is required")
		}
	}
}
