// config/config.go
package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Config holds the entire application configuration
type Config struct {
	App          AppConfig            `yaml:"app"`
	ProductDB    DatabaseConfig       `yaml:"product_database"` // Renamed to clarify its purpose
	StockDB      DatabaseConfig       `yaml:"stock_database"`   // Renamed to clarify its purpose
	Environments map[string]EnvConfig `yaml:"environments"`     // Renamed for clarity
}

// AppConfig holds application-specific configurations
type AppConfig struct {
	Port int `yaml:"port"`
}

// DatabaseConfig holds database connection details
type DatabaseConfig struct {
	Driver string `yaml:"driver"` // e.g., sqlite, postgres
	Source string `yaml:"source"` // Database source string
}

// EnvConfig holds environment-specific configurations
type EnvConfig struct {
	ProductDB DatabaseConfig `yaml:"product_database"`
	StockDB   DatabaseConfig `yaml:"stock_database"`
}

// LoadConfig loads the configuration based on the provided environment
func LoadConfig(env string) (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	// If environment is specified, override default settings
	if env != "" {
		if envConfig, exists := cfg.Environments[env]; exists {
			cfg.ProductDB = envConfig.ProductDB
			cfg.StockDB = envConfig.StockDB
		}
	}

	return &cfg, nil
}
