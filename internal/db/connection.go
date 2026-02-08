package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alexvitayu/gorm-project/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	maxOpenConns    = 10
	maxIdleConns    = 5
	connMaxLifetime = time.Hour
)

// Open creates a GORM connection with logging and pool settings.
func Open(cfg config.DBConfig) (*gorm.DB, error) {
	if cfg.DSN == "" {
		return nil, errors.New("DATABASE_DSN is empty")
	}

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("fail to connect to DB: %w", err)
	}

	pgDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	if err = pgDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Ping to DB is successful!")

	pgDB.SetMaxOpenConns(maxOpenConns)
	pgDB.SetMaxIdleConns(maxIdleConns)
	pgDB.SetConnMaxLifetime(connMaxLifetime)

	return db, nil
}
