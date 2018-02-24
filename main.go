package main

import (
	"fmt"
	"go-simple-server/server"
)

func main() {
	fmt.Println("Start server ...")
	server.InitServer()
	fmt.Println("Server started")
}
