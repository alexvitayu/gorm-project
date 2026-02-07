package models

import "time"

type Movie struct {
	ID          uint
	Title       string `gorm:"size:150;not null"`
	Genre       string
	ReleasedAt  time.Time
	Description string
}
