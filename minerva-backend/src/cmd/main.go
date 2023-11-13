package main

import (
	"github.com/Code-Mozart/minerva/minerva-backend/internal/server"
)

func main() {
	server := &server.Server{}
	server.Init()
	server.Run()
}
