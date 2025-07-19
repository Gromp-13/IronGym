package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Gromp-13/IronGym/internal/db/repository"
	"github.com/joho/godotenv"
)

type User struct {
	ID          int32
	LastName    string
	FirstName   string
	MiddleName  string
	PhoneNumber string
	BirthDate   time.Time
	CardBarcode string
}

func Conectdb() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		log.Fatal("DB_URL не задан")
	}

	bd, err = repository.New(DBURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := bd.GetClients()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range data {
		fmt.Printf("%d: %s, lastname: %d, firstname: %d, middlename: %d, phonenumber: %d, birthdate: %d, cardbarcode: %d",
			item.ID, item.LastName, item.FirstName, item.MiddleName, item.PhoneNumber, item.BirthDate, item.CardBarcode)
	}

}
