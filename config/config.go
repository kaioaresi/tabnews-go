package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DBCredentials struct {
	Environment       string `mapstructure:"environment"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
}

func NewConfig() (*DBCredentials, error) {
	creds, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return creds, nil
}

func LoadConfig() (*DBCredentials, error) {
	viper.SetConfigFile(".env.development")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("error config file not found!: %v", err)
		}
		return nil, fmt.Errorf("error to load environments from file: %v", err)
	}

	var creds DBCredentials
	viper.Unmarshal(&creds)

	return &creds, nil
}

func (d *DBCredentials) StringConnection() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/tabnews?sslmode=disable", d.POSTGRES_USER, d.POSTGRES_PASSWORD, d.POSTGRES_HOST, d.POSTGRES_PORT)
}
