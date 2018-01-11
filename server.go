package main

import (
	"github.com/NYTimes/marvin"
	"google.golang.org/appengine"
)

func main() {
	marvin.Init(NewService())
	appengine.Main()
}
