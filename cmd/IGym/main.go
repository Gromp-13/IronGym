package main

import (
	"github.com/Gromp-13/IronGym/internal/config"
	"github.com/Gromp-13/IronGym/internal/db"
)

func main() {

	db.Connectdb()

	defer db.RepoClose()

	config.Config()

}
