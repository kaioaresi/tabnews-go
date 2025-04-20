package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DBCredentials struct {
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
}

func NewConfig() (*DBCredentials, error) {
	creds, err := LoadEnviroments(".")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return creds, nil
}

func LoadEnviroments(path string) (*DBCredentials, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env.development")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Error to load environments from file: %v", err)
	}

	var creds DBCredentials
	viper.Unmarshal(&creds)

	return &creds, nil
}
