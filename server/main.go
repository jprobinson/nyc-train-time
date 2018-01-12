package main

import (
	"github.com/jprobinson/dialogflow"
	nyctraintime "github.com/jprobinson/nyc-train-time"
	"google.golang.org/appengine"
)

func main() {
	dialogflow.Init(nyctraintime.NewGoogleService(), nyctraintime.GoodbyeMiddleware)
	appengine.Main()
}
