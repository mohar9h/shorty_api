package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig
	Postgres  PostgresConfig
	Redis     RedisConfig
	Cors      CorsConfig
	Logger    LoggerConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type CorsConfig struct {
	AllowOrigins string
}

type PostgresConfig struct {
	Host            string
	Port            string
	Username        string
	Password        string
	Database        string
	SslMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Database           string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
}

type RateLimitConfig struct {
	RequestsPerSecond int
	BurstSize         int
}

func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(configPath, "yml")
	if err != nil {
		log.Fatalf("LoadConfig failed: %v", err)
	}
	config, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("ParseConfig failed: %v", err)
	}

	return config
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	switch env {
	case "docker":
		return "./config/docker"
	case "production":
		return "./config/production"
	default:
		return "./config/development"
	}
}
