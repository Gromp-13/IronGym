package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Conectdb() {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:pass@localhost:5432/irongum")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	_ = conn

}
