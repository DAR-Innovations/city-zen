package config

type ServerConfig struct {
	Port      int    `mapstructure:"SERVER_PORT"`
	ClientURL string `mapstructure:"CLIENT_URL"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`
}
