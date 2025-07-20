package main

import (
	"github.com/Gromp-13/IronGym/internal/config"
	"github.com/Gromp-13/IronGym/internal/db"
	"github.com/Gromp-13/IronGym/internal/db/repository"
)

var repo *repository.PGRepo

func main() {

	repo, _ = db.Connectdb()
	defer repo.Close()

	config.Config()

}
