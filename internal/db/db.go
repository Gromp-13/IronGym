package db

import (
	"log"
	"os"

	"github.com/Gromp-13/IronGym/internal/db/repository"
	"github.com/Gromp-13/IronGym/internal/models"
	"github.com/joho/godotenv"
)

func Connectdb() []models.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		log.Fatal("DB_URL не задан")
	}

	repo, err := repository.New(DBURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	clients, err := repo.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	return clients

}
