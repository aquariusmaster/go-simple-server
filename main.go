package main

import (
	"fmt"
	"github.com/aquariusmaster/go-simple-server/server"
)

func main() {
	fmt.Println("Start server ...")
	server.InitServer()
	fmt.Println("Server started")
}
