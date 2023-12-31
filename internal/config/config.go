package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string `yaml:"env" env-default:"local"`
	Storage_name string `yaml:"storage_name"`
	HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8082"`
	Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
	Idle_timeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User         string        `yaml:"user" env-required:"true"`
	Password     string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() *Config {
	configPath := "../../config/config.yaml"

	fmt.Println(configPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
