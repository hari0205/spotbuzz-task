package setup

import (
	"os"

	"github.com/hari0205/spotbuzz-task/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func SetUpDB() {
	var err error
	if err = godotenv.Load(); err != nil {
		panic(err)
	}
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default, TranslateError: true})
	if err != nil {
		panic(err.Error())
	}
	database.Debug().AutoMigrate(&models.Player{})

	GormDB = database
}
