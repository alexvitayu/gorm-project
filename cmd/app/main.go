package main

import (
	"log"
	"os"
	"time"

	"github.com/alexvitayu/gorm-project/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	if err := godotenv.Load(".env.development"); err != nil {
		log.Fatalf("переменные окружения не загружены: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("ошибка подключения: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("ошибка доступа к пулу: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("ошибка пинга базы: %v", err)
	}

	log.Println("Пинг базы данных прошёл успешно!")

	//if err := db.Create(&User{Name: "Анна", Email: "anna@example.com"}).Error; err != nil {
	//	log.Fatalf("ошибка вставки: %v", err)
	//}

	//var user User
	//if err := db.First(&user).Error; err != nil {
	//	log.Fatalf("ошибка чтения: %v", err)
	//}
	//
	//log.Printf("пользователь загружен: %s <%s>", user.Name, user.Email)

	var movie models.Movie
	if err = db.First(&movie).Error; err != nil {
		log.Fatalf("ошибка чтения: %v", err)
	}

	log.Printf("фильм загружен: %s <%s>", movie.Title, movie.Genre)
}
