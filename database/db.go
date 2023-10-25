package database

import (
	"api-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	dsn := "host=localhost user=admin password=bcd127 dbname=personalidades port=5432 sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco")
	}

	DB.AutoMigrate(&models.Aluno{})
}
