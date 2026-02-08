package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Movie struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:150;not null"`
	Genre       string
	ReleasedAt  time.Time
	Description string
	Rating      *decimal.Decimal `gorm:"type:numeric(3,1)"`

	// Внешний ключ для Director
	DirectorID uint
	Director   Director `gorm:"foreignKey:DirectorID"`

	// Связь many2many с Actor
	Actors []Actor `gorm:"many2many:movie_actors"`
}
