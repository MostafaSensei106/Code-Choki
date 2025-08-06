package config

import (
	"fmt"
	"os"
	"path/filepath"
	//"runtime"

	"gopkg.in/yaml.v3"
)

type Config struct{}

func DefaultConfig() *Config {
	return &Config{}
}
func getConfigDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".CodeChoki/config")
}

func LoadConfig() (*Config, error) {
	configDirectory := getConfigDirectory()
	configPath := filepath.Join(configDirectory, "config.yaml")

	//create config directory if it doesn't exist
	if err := os.MkdirAll(configDirectory, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	// if config file doesn't exist, create it
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := DefaultConfig()
		if err := defaultConfig.Save(); err != nil {
			return nil, fmt.Errorf("failed to save default config: %v", err)
		}
		return defaultConfig, nil
	}

	//load existing config

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}
	return &conf, nil
}

func (c *Config) Save() error {
	configDirctory := getConfigDirectory()
	configPath := filepath.Join(configDirctory, "config.yaml")
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}
	return nil
}
