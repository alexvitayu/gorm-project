package db

import (
	"fmt"

	"github.com/alexvitayu/gorm-project/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Movie{}, &models.Actor{}, &models.Director{}); err != nil {
		return fmt.Errorf("fail to migrate: %w", err)
	}
	return nil
}
