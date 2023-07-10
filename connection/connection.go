package connection

import (
	"keuangan/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Can't Load Env")
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed Connect Database")
	}

	db.AutoMigrate(&entity.Pendaftaran{}, &entity.Pembangunan{}, &entity.Semesteran{}, &entity.Sidang{}, &entity.Wisuda{}, &entity.LapBulanan{}, &entity.LapTahunan{})
	return db
}
