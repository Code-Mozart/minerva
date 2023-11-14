package main

import (
	"flag"
	"fmt"

	"github.com/Code-Mozart/minerva/minerva-backend/internal/server"
)

func main() {
	server := &server.Server{}
	server.Init()

	port := flag.String("port", "8134", "Port to listen on (default is 8134 as in 'm1n3rv4')")
	flag.Parse()
	if port == nil {
		panic("Port is nil")
	}

	address := fmt.Sprintf(":%s", *port)
	server.Run(address)
}
