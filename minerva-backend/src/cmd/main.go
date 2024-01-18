package main

import (
	"fmt"
	"log"

	"github.com/Code-Mozart/minerva/minerva-backend/initializers"
	"github.com/Code-Mozart/minerva/minerva-backend/internal/server"
)

const DEFAULT_PORT string = "8134"

func main() {
	env, err := initializers.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.Initialize(env)
	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%s", env.Port)
	server.Run(address)
}
