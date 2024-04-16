package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string `envconfig:"ORDER_DB_HOST"`
	Port     string `envconfig:"ORDER_DB_PORT"`
	User     string `envconfig:"ORDER_DB_USER"`
	Password string `envconfig:"ORDER_DB_PASSWORD"`
	Name     string `envconfig:"ORDER_DB_NAME"`
}

type JwtConfig struct {
	Key string `envconfig:"ORDER_JWT_SECRET"`
}

type AppConfig struct {
	Database DatabaseConfig
	Jwt      JwtConfig
}

var Cfg = Init()

func Init() *AppConfig {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}

func NewConfig() (*AppConfig, error) {
	var cfg AppConfig
	if err := envconfig.Process("ORDER", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
