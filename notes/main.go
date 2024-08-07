package main

import (
	"github.com/vyacheslavprod/microservices/notes/server"
)

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
	
}