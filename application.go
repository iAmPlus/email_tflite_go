package main

import (
	"flag"
	"github.com/iAmPlus/skills-text-extraction-go/app"
	"log"
)

func main() {

	a := app.CreateApplication()
	var port = flag.Int("port", 8080, "port on which the server should run")

	err := a.StartServer(*port)
	if err != nil {
		log.Println("unable to start server")
	}
}
