package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Wechat WechatConfig
	Server ServerConfig
}

type WechatConfig struct {
	AppID         string `mapstructure:"app_id"`
	AppSecret     string `mapstructure:"app_secret"`
	Token         string
	EncodingAESKey string `mapstructure:"encoding_aes_key"`
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*Configuration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}