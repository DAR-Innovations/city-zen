package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfg *Config

type Config struct {
	Environment  string `mapstructure:"ENVIRONMENT"`
	IsDebug      bool   `mapstructure:"DEBUG"`
	UploadFolder string `mapstructure:"UPLOAD_FOLDER"`

	Database DatabaseConfig `mapstructure:",squash"`
	Server   ServerConfig   `mapstructure:",squash"`
	JWT      JWTConfig      `mapstructure:",squash"`
	AI       AIConfig       `mapstructure:",squash"`
}

func LoadConfig(file string) (*Config, error) {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read config file: %v", err)
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func GetConfig() *Config {
	if cfg == nil {
		log.Fatal("Config not loaded")
	}
	return cfg
}
