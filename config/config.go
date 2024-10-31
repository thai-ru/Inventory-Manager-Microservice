// config/config.go
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App      AppSettings    `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}

type AppSettings struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

func LoadConfig(env string) (*AppConfig, error) {
	config := &AppConfig{}

	// Load .env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Override with environment variables
	if driver := os.Getenv("DB_DRIVER"); driver != "" {
		config.Database.Driver = driver
	} else {
		config.Database.Driver = "sqlite" // Default value
	}

	if source := os.Getenv("DB_SOURCE"); source != "" {
		config.Database.Source = source
	} else {
		config.Database.Source = "file:dev.db?cache=shared&mode=rwc" // Default value
	}

	if port := os.Getenv("APP_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.App.Port = p
		} else {
			return nil, fmt.Errorf("invalid APP_PORT: %s", port)
		}
	} else {
		config.App.Port = 50051 // Default port
	}

	return config, nil
}
