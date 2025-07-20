package db

import (
	"log"
	"os"

	"github.com/Gromp-13/IronGym/internal/db/repository"
	"github.com/Gromp-13/IronGym/internal/models"
	"github.com/joho/godotenv"
)

var Repo *repository.PGRepo

func Connectdb() []models.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		log.Fatal("DB_URL не задан")
	}

	var errRepo error
	Repo, errRepo = repository.New(DBURL)
	if errRepo != nil {
		log.Fatal(errRepo.Error())
	}

	clients, err := Repo.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	return clients

}

func RepoClose() {
	if Repo != nil {
		Repo.Close()
	}
}
