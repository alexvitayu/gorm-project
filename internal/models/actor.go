package models

type Actor struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`

	// Связь many2many с Movie
	Movies []Movie `gorm:"many2many:movie_actors"`
}
