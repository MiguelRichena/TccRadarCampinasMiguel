package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


tpye Config struct {
	DBHost     string
	DBPort     string 
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBTimezone string

}

func Load() (*Config, error) {
	// carrega .env em dev

	_ = godotenv.Load()

	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBSSLMode: os.Getenv("DB_SSLMODE"),
		DBTimezone: os.Getenv("DB_TIMEZONE"),
	}

	if cfg.DB_HOST == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("Variaveis de ambiente de DB não configuradas")
	}
	return cfg, nil
}