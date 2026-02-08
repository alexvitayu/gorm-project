package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string
}

func Load() *DBConfig {
	if err := godotenv.Load(".env.development"); err != nil {
		log.Fatalf("env variables are not loaded: %v", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	return &DBConfig{DSN: dsn}
}
