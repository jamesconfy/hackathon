package config

import "github.com/spf13/viper"

type Config struct {
	MODE               string `mapstructure:"MODE"`
	ADDR               string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN   string `mapstructure:"SECRET_KEY_TOKEN"`
	POSTGRES_HOST      string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_USERNAME  string `mapstructure:"POSTGRES_USERNAME"`
	POSTGRES_PASSWORD  string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DBNAME    string `mapstructure:"POSTGRES_DBNAME"`
	POSTGRES_PORT      string `mapstructure:"POSTGRES_PORT"`
	AWS_ACCESS_KEY     string `mapstructure:"AWS_ACCESS_KEY"`
	AWS_SECRET_KEY     string `mapstructure:"AWS_SECRET_KEY"`
	DEFAULT_USER_ID    string `mapstructure:"DEFAULT_USER_ID"`
	DEFAULT_ACCOUNT_ID string `mapstructure:"DEFAULT_ACCOUNT_ID"`
}

var AppConfig *Config

func init() {
	viper.AutomaticEnv()

	AppConfig = &Config{
		MODE:               viper.GetString("MODE"),
		ADDR:               viper.GetString("ADDR"),
		SECRET_KEY_TOKEN:   viper.GetString("SECRET_KEY_TOKEN"),
		POSTGRES_HOST:      viper.GetString("POSTGRES_HOST"),
		POSTGRES_USERNAME:  viper.GetString("POSTGRES_USERNAME"),
		POSTGRES_PASSWORD:  viper.GetString("POSTGRES_PASSWORD"),
		POSTGRES_DBNAME:    viper.GetString("POSTGRES_DBNAME"),
		POSTGRES_PORT:      viper.GetString("POSTGRES_PORT"),
		AWS_ACCESS_KEY:     viper.GetString("AWS_ACCESS_KEY"),
		AWS_SECRET_KEY:     viper.GetString("AWS_SECRET_KEY"),
		DEFAULT_USER_ID:    viper.GetString("DEFAULT_USER_ID"),
		DEFAULT_ACCOUNT_ID: viper.GetString("DEFAULT_ACCOUNT_ID"),
	}
}
