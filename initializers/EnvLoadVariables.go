package initializers

import (
	"log"

	"github.com/subosito/gotenv"
)

func EnvLoadVariables() {

	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
