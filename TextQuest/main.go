package main

import (
	"log"
	"TextQuest/server"
)
func main()  {
	log.Fatal(server.RunHTTPServer(":8080"))
}
