package main

import (
	"log"
	"os"

	"github.com/jprobinson/dialogflow"
	nyctraintime "github.com/jprobinson/nyc-train-time"
)

func main() {
	svc, err := nyctraintime.NewService()
	if err != nil {
		log.Fatal(err)
	}

	err = dialogflow.Run(svc, os.Getenv("AUDIENCE"))
	if err != nil {
		log.Fatal(err)
	}
}
