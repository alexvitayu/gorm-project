package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Movie struct {
	ID          uint
	Title       string `gorm:"size:150;not null"`
	Genre       string
	ReleasedAt  time.Time
	Description string
	Rating      *decimal.Decimal `gorm:"type:numeric(3,1)"`
}
