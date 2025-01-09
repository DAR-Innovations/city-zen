package config

type AIConfig struct {
	ModelPath    string `mapstructure:"AI_MODEL_PATH"`
	ConfigPath   string `mapstructure:"AI_CONFIG_PATH"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`
}
