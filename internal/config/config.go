package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	AppDir         = "shihtzu"
	ConfigFileName = "config.yaml"
)

type ConfigParser struct{}

// LoadConfig loads the application configuration using Viper
func LoadConfig() error {
	fmt.Println("Starting to load config...") // Debug
	parser := ConfigParser{}
	configPath, err := parser.getConfigFileOrCreateIfMissing()
	if err != nil {
		return err
	}
	fmt.Printf("Config file path %s\n", *configPath)
	viper.SetConfigFile(*configPath) // Use the config path found or created
	viper.AutomaticEnv()             // Automatically read in environment variables

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Error reading config file: %w", err)
	}

	return nil
}

// getConfigFileOrCreateIfMissing returns the config file path or creates it if it doesn't exist
func (parser ConfigParser) getConfigFileOrCreateIfMissing() (*string, error) {
	configDir := os.Getenv("XDG_CONFIG_HOME")

	// If XDG_CONFIG_HOME is not set, use ~/.config as fallback
	if configDir == "" {
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	shihtzuConfigDir := filepath.Join(configDir, AppDir)
	err := os.MkdirAll(shihtzuConfigDir, os.ModePerm) // Create the directory if it doesn't exist
	if err != nil {
		return nil, fmt.Errorf("Error creating config directory: %w", err)
	}

	configFilePath := filepath.Join(shihtzuConfigDir, ConfigFileName)
	err = parser.createConfigFileIfMissing(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error creating config file: %w", err)
	}

	return &configFilePath, nil
}

// createConfigFileIfMissing creates the config file with default settings if it doesn't exist
func (parser ConfigParser) createConfigFileIfMissing(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write default configuration settings
		file.WriteString("settings:\n  enable_logging: false\n  enable_mousewheel: true\n")
	}
	return nil
}
