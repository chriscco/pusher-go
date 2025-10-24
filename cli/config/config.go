package config

import (
	"github.com/spf13/viper"
)

type Model struct {
	ApiKey    string `mapstructure:"api_key"`
	ModelName string `mapstructure:"model_name"`
}

type Email struct {
	To       []string `mapstructure:"to"`
	From     string   `mapstructure:"from"`
	Subject  string   `mapstructure:"subject"`
	Password string   `mapstructure:"password"`
}

type GNews struct {
	ApiKey     string   `mapstructure:"api_key"`
	Endpoint   string   `mapstructure:"endpoint"`
	Categories []string `mapstructure:"categories"`
	Languages  []string `mapstructure:"languages"`
	Countries  []string `mapstructure:"countries"`
	MaxResults int      `mapstructure:"max_results"`
}

type Config struct {
	Model Model `mapstructure:"model"`
	Email Email `mapstructure:"email"`
	GNews GNews `mapstructure:"g_news"`
}

func LoadConfig(filename string) (*Config, error) {
	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
