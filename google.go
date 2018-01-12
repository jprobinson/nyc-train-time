package nyctraintime

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/NYTimes/marvin"
	"github.com/go-kit/kit/endpoint"
	"github.com/jprobinson/dialogflow"
)

type googleService struct {
	key string
}

func NewGoogleService() dialogflow.GoogleActionService {
	return &googleService{key: os.Getenv("MTA_KEY")}
}

func (g *googleService) Actions() map[string]dialogflow.GoogleActionHandler {
	return map[string]dialogflow.GoogleActionHandler{
		"my_next_train_request":      g.myTrain,
		"my_following_train_request": g.myFollowingTrain,
		"save_my_stop_request":       g.saveMyStopAction,
		"next_train_request":         g.nextTrain,
		"following_train_request":    g.followingTrain,
	}
}

// GoodbyeMiddleware will add a generic, random "good bye" message to the end of eligible
// responses.
func GoodbyeMiddleware(ep endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		//call our action
		re, err := ep(ctx, r)
		if err != nil {
			return re, err
		}

		// if no error, add a generic goodbye
		switch res := re.(type) {
		case *dialogflow.GoogleFulfillmentResponse:
			if res.Speech == "" {
				return res, err
			}
			bye := " ..." +
				goodbyes[rand.New(rand.NewSource(time.Now().Unix())).Intn(len(goodbyes)-1)]
			res.Speech = res.Speech + bye
			res.DisplayText = res.Speech
			return res, err
		default:
			return res, err
		}
	}
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

func (g *googleService) myTrain(ctx context.Context, r *dialogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	uid := r.OriginalRequest.Data.User.UserID
	if uid == "" {
		return simpleGoogleResponse("Sorry, you need to be logged in for that to work")
	}
	mys, err := getMyStop(ctx, uid)
	if err == datastore.ErrNoSuchEntity {
		return simpleGoogleResponse(
			"It looks like you haven't saved your personalized subway stop yet! Ask NYC Train Time to \"save my stop\" to create or update your stop.")
	}
	if err != nil {
		log.Debugf(ctx, "unable to get my stop: %s", err)
		return simpleGoogleResponse(
			"Sorry, we were unable to access the train feed. Please try again")
	}

	ft, err := parseFeed(mys.Line)
	if err != nil {
		log.Debugf(ctx, "unable to parse line: %s", mys.Line)
		return simpleGoogleResponse(
			fmt.Sprintf("sorry, the %s line is not available yet", mys.Line))
	}

	return simpleGoogleResponse(
		g.getNextTrainDialog(ctx, ft, mys.Line, mys.Stop, mys.Dir))
}

func (g *googleService) myFollowingTrain(ctx context.Context, r *dialogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	uid := r.OriginalRequest.Data.User.UserID
	if uid == "" {
		return simpleGoogleResponse("Sorry, you need to be logged in for that to work")
	}
	mys, err := getMyStop(ctx, uid)
	if err == datastore.ErrNoSuchEntity {
		return simpleGoogleResponse(
			"It looks like you haven't saved your personalized subway stop yet! Ask NYC Train Time to \"save my stop\" to create or update your stop.")
	}
	if err != nil {
		log.Debugf(ctx, "unable to get my stop: %s", err)
		return simpleGoogleResponse(
			"Sorry, we were unable to access the train feed. Please try again")
	}

	ft, err := parseFeed(mys.Line)
	if err != nil {
		log.Debugf(ctx, "unable to parse line: %s", mys.Line)
		return simpleGoogleResponse(
			fmt.Sprintf("sorry, the %s line is not available yet", mys.Line))
	}

	return simpleGoogleResponse(
		g.getFollowingTrainDialog(ctx, ft, mys.Line, mys.Stop, mys.Dir))
}

func (g *googleService) saveMyStopAction(ctx context.Context, r *dialogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	uid := r.OriginalRequest.Data.User.UserID
	if uid == "" {
		return simpleGoogleResponse("Sorry, you need to be logged in for that to work")
	}

	line := strings.ToUpper(r.Result.Parameters["subway-line"].(string))
	stop := r.Result.Parameters["subway-stop"].(string)
	dir := r.Result.Parameters["subway-direction"].(string)

	err := saveMyStop(ctx, uid, line, stop, dir)
	if err != nil {
		return nil, marvin.NewJSONStatusResponse(map[string]string{
			"error": "unable to complete request: " + err.Error(),
		}, http.StatusInternalServerError)
	}
	return simpleGoogleResponse(fmt.Sprintf(
		"Successfully saved your stop, %s bound %s trains at %s. To update your stop again, ask NYC Train Time to \"save my stop\". ",
		dir, line, stop))
}

func (g *googleService) nextTrain(ctx context.Context, r *dialogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	line := strings.ToUpper(r.Result.Parameters["subway-line"].(string))
	stop := r.Result.Parameters["subway-stop"].(string)
	dir := r.Result.Parameters["subway-direction"].(string)

	ft, err := parseFeed(line)
	if err != nil {
		log.Debugf(ctx, "unable to parse line: %s", line)
		return simpleGoogleResponse(fmt.Sprintf("Sorry, the %s line is not available yet", line))
	}

	return simpleGoogleResponse(g.getNextTrainDialog(ctx, ft, line, stop, dir) +
		" If you would like me to remember your stop, ask NYC Train Time to \"save my stop\" and then ask for MY stop next time. ")
}

func (g *googleService) followingTrain(ctx context.Context, r *dialogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	line := strings.ToUpper(r.Result.Parameters["subway-line"].(string))
	stop := r.Result.Parameters["subway-stop"].(string)
	dir := r.Result.Parameters["subway-direction"].(string)
	ft, err := parseFeed(line)
	if err != nil {
		log.Debugf(ctx, "unable to parse line: %s", line)
		return simpleGoogleResponse(
			fmt.Sprintf("Sorry, the %s line is not available yet.", line))
	}
	return simpleGoogleResponse(g.getFollowingTrainDialog(ctx, ft, line, stop, dir))
}

const source = "Where's The Train (NYC)"

func simpleGoogleResponse(res string) (*dialogflow.GoogleFulfillmentResponse, error) {
	return &dialogflow.GoogleFulfillmentResponse{
		Speech:      res,
		DisplayText: res,
		Source:      source,
	}, nil
}
