package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DSN string `mapstructure:"DB_DSN"`

	DB_CFG struct {
		MaxOpenConn     int32         `mapstructure:"maxOpenConn"`
		MaxIdleConn     int32         `mapstructure:"maxIdleConn"`
		ConnMaxLifeTime time.Duration `mapstructure:"connMaxLifeTime"`
	} `mapstructure:"postgres"`

	Redis struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Pass string `mapstructure:"pass"`
	} `mapstructure:"redis"`
}

func Load(path string) (*Config, error) {
	var config Config

	readFromConfigFile()
	readFromEnvFile(path)

	viper.AutomaticEnv()

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &config, nil
}

func readFromConfigFile() {

	viper.SetConfigFile("./config/cfg.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func readFromEnvFile(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading env file, %s", err)
	}
}
