package models

type Director struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`

	// Обратная связь one2many с Movie
	Movies []Movie `gorm:"foreignKey:DirectorID"`
}
