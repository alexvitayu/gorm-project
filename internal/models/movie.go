package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Movie struct {
	ID          uint
	Title       string `gorm:"size:150;not null"`
	Genre       string
	ReleasedAt  time.Time
	Description string
	Rating      decimal.Decimal `gorm:"type:numeric(3,1)"`
}
