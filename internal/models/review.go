package models

type Review struct {
	ID      uint  `gorm:"primaryKey"`
	MovieID uint  // внешний ключ на movies.id
	Movie   Movie `gorm:"foreignKey:MovieID"` // belongs to: ревью принадлежит фильму
	Text    string
	Score   int
}
