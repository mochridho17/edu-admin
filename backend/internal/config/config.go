package config

import (
	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	JWTExpiry  int    `mapstructure:"JWT_EXPIRY"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "edu-admin-secret-key-2025"
	}
	if cfg.JWTExpiry == 0 {
		cfg.JWTExpiry = 24
	}

	AppConfig = &cfg
	return &cfg, nil
}

func (c *Config) DSN() string {
	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=" + c.DBSSLMode
}
