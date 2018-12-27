package main

import (
	"log"

	"github.com/jprobinson/dialogflow"
	nyctraintime "github.com/jprobinson/nyc-train-time"
)

func main() {
	svc, err := nyctraintime.NewService()
	if err != nil {
		log.Fatal(err)
	}

	err = dialogflow.Run(svc)
	if err != nil {
		log.Fatal(err)
	}
}
