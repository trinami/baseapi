package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBType     string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	DBSSLMode  string
	DBTimeZone string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	viper.SetDefault("DB_TYPE", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "testdb")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("DB_TimeZone", "UTC")

	config := &Config{
		DBType:     viper.GetString("DB_TYPE"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBSSLMode:  viper.GetString("DB_SSLMode"),
		DBTimeZone: viper.GetString("DB_TimeZone"),
	}

	return config
}
