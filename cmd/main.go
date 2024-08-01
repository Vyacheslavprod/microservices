package main

import (
	"github.com/vyacheslavprod/microservices/server"
)

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
	
}