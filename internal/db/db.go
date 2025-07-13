package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func Conectdb() {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:@localhost:5432/irongumdb")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(), "SELECT * FROM clients;")
	fmt.Println(row)

}
