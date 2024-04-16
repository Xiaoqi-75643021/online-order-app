package config_test

import (
	"fmt"
	"testing"

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

func TestLoadConfig(t *testing.T) {
	var cfg AppConfig
	err := envconfig.Process("ORDER", &cfg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cfg)
}
