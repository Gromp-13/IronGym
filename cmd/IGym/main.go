package main

import (
	"github.com/Gromp-13/IronGym/internal/config"
	"github.com/Gromp-13/IronGym/internal/ui/screens"
)

func main() {

	p := 0
	if p == 5 {
		screens.NewClientScreen()
	}

	config.ConfScreen()

}
