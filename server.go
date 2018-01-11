package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/jprobinson/dialogflow"
	"google.golang.org/appengine"
)

func main() {
	dialogflow.Run(NewGoogleServer(), func(ep endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, r interface{}) (interface{}, error) {
			//call our action
			re, err := ep(ctx, r)

			switch res := re.(type) {
			// add generic goodbye suffix to response
			case *dialogflow.GoogleFulfillmentResponse:
				bye := " ..." + goodbyes[rand.New(rand.NewSource(time.Now().Unix())).Intn(len(goodbyes)-1)]
				if res.Speech == "" {
					return res, err
				}
				res.Speech = res.Speech + bye
				res.DisplayText = res.Speech
				return res, err
			default:
				return res, err
			}
		}
	})
	appengine.Main()
}

var goodbyes = []string{
	"Ok, bye!",
	"Bye bye now",
	"Peace out!",
	"Goodbye",
	"Hope you can catch the train!",
	"Hope you can make it!",
	"Adios!",
	"Au revoir",
	"Have a good trip!",
	"Have a good ride!",
	"Have a save trip!",
	"Save travels!",
}
