package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file '%s' does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can't read config: %s", err)
	}

	return &cfg
}
