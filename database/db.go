package database

import (
	"fmt"
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("DBPORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	DB.AutoMigrate(&models.Aluno{})
}
