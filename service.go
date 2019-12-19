package nyctraintime

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/NYTimes/marvin"
	"github.com/go-kit/kit/endpoint"
	"github.com/jprobinson/dialogflow"
)

type service struct {
	key string
	db  *db
	hc  *http.Client
}

func NewService() (dialogflow.FulfillmentService, error) {
	ctx := context.Background()

	db, err := newDB(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return nil, err
	}
	return &service{
		key: os.Getenv("MTA_KEY"),
		db:  db,
		hc:  &http.Client{Timeout: 2 * time.Second},
	}, nil
}

func (g *service) Intents() map[string]dialogflow.IntentHandler {
	return map[string]dialogflow.IntentHandler{
		"My Next Train":      g.myTrain,
		"My Following Train": g.myFollowingTrain,
		"Save My Stop":       g.saveMyStopAction,
		"Next Train":         g.nextTrain,
		"Following Train":    g.followingTrain,
	}
}

// Middleware will add a generic, random "good bye" message to the end of eligible
// responses.
func (s *service) Middleware(ep endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		//call our action
		re, err := ep(ctx, r)
		if err != nil {
			kit.LogErrorMsg(ctx, err, "encountered an error")
			return re, err
		}

		// if no error, add a generic goodbye
		switch res := re.(type) {
		case *dialogflow.FulfillmentResponse:
			if res.FulfillmentText == "" {
				return res, err
			}
			bye := " ..." +
				goodbyes[rand.New(rand.NewSource(time.Now().Unix())).Intn(len(goodbyes)-1)]
			res.FulfillmentText = res.FulfillmentText + bye
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
	"Toodles!",
	"Until next time!",
}

func (s *service) myTrain(ctx context.Context, r *dialogflow.Request) (*dialogflow.FulfillmentResponse, error) {
	uid := r.OriginalDetectIntentRequest.Payload.User.UserID
	if uid == "" {
		return simpleResponse(
			"It looks like you aren't logged in yet. Please link your account with the Google Assistant.")
	}
	mys, err := s.db.getMyStop(ctx, uid)
	if err == datastore.ErrNoSuchEntity {
		return &dialogflow.FulfillmentResponse{
			FollowupEventInput: &dialogflow.EventInput{Name: "save_stop"},
			Source:             source,
		}, nil
	}
	if err != nil {
		kit.LogErrorMsg(ctx, err, "unable to get my stops")
		return simpleResponse(
			"Sorry, we were unable to access the train feed. Please try again")
	}

	ft, err := parseFeed(mys.Line)
	if err != nil {
		kit.LogMsg(ctx, "unable to parse line: "+mys.Line)
		return simpleResponse(
			fmt.Sprintf("sorry, the %s line is not available yet", mys.Line))
	}

	return simpleResponse(
		s.getNextTrainDialog(ctx, ft, mys.Line, mys.Stop, mys.Dir))
}

func (s *service) myFollowingTrain(ctx context.Context, r *dialogflow.Request) (*dialogflow.FulfillmentResponse, error) {
	uid := r.OriginalDetectIntentRequest.Payload.User.UserID
	if uid == "" {
		return simpleResponse("Sorry, you need to be logged in for that to work")
	}
	mys, err := s.db.getMyStop(ctx, uid)
	if err == datastore.ErrNoSuchEntity {
		return &dialogflow.FulfillmentResponse{
			FollowupEventInput: &dialogflow.EventInput{Name: "save_stop"},
			Source:             source,
		}, nil
	}
	if err != nil {
		kit.LogErrorMsg(ctx, err, "unable to get my stop")
		return simpleResponse(
			"Sorry, we were unable to access the train feed. Please try again")
	}

	ft, err := parseFeed(mys.Line)
	if err != nil {
		kit.LogMsg(ctx, "unable to parse line: "+mys.Line)
		return simpleResponse(
			fmt.Sprintf("sorry, the %s line is not available yet", mys.Line))
	}

	return simpleResponse(
		s.getFollowingTrainDialog(ctx, ft, mys.Line, mys.Stop, mys.Dir))
}

func (s *service) saveMyStopAction(ctx context.Context, r *dialogflow.Request) (*dialogflow.FulfillmentResponse, error) {
	uid := r.OriginalDetectIntentRequest.Payload.User.UserID
	if uid == "" {
		return simpleResponse("Sorry, you need to be logged in for that to work")
	}

	line := strings.ToUpper(r.QueryResult.Parameters["subway-line"].(string))
	stop := r.QueryResult.Parameters["subway-stop"].(string)
	dir := r.QueryResult.Parameters["subway-direction"].(string)

	err := s.db.saveMyStop(ctx, uid, line, stop, dir)
	if err != nil {
		return nil, marvin.NewJSONStatusResponse(map[string]string{
			"error": "unable to complete request: " + err.Error(),
		}, http.StatusInternalServerError)
	}

	ft, err := parseFeed(line)
	if err != nil {
		kit.LogMsg(ctx, "unable to parse line: "+line)
		return simpleResponse(
			fmt.Sprintf("sorry, the %s line is not available yet", line))
	}

	return simpleResponse("Successfully saved your stop. " +
		s.getNextTrainDialog(ctx, ft, line, stop, dir))
}

func (g *service) nextTrain(ctx context.Context, r *dialogflow.Request) (*dialogflow.FulfillmentResponse, error) {
	line := strings.ToUpper(r.QueryResult.Parameters["subway-line"].(string))
	stop := r.QueryResult.Parameters["subway-stop"].(string)
	dir := r.QueryResult.Parameters["subway-direction"].(string)

	ft, err := parseFeed(line)
	if err != nil {
		kit.LogMsg(ctx, "unable to parse line: "+line)
		return simpleResponse(fmt.Sprintf("Sorry, the %s line is not available yet", line))
	}

	return simpleResponse(g.getNextTrainDialog(ctx, ft, line, stop, dir) +
		" If you would like me to remember your stop, ask NYC Train Time to \"save my stop\" and then ask for MY stop next time. ")
}

func (g *service) followingTrain(ctx context.Context, r *dialogflow.Request) (*dialogflow.FulfillmentResponse, error) {
	line := strings.ToUpper(r.QueryResult.Parameters["subway-line"].(string))
	stop := r.QueryResult.Parameters["subway-stop"].(string)
	dir := r.QueryResult.Parameters["subway-direction"].(string)
	ft, err := parseFeed(line)
	if err != nil {
		kit.LogMsg(ctx, "unable to parse line: "+line)
		return simpleResponse(
			fmt.Sprintf("Sorry, the %s line is not available yet.", line))
	}
	return simpleResponse(g.getFollowingTrainDialog(ctx, ft, line, stop, dir))
}

const source = "Where's The Train (NYC)"

func simpleResponse(res string) (*dialogflow.FulfillmentResponse, error) {
	return &dialogflow.FulfillmentResponse{
		FulfillmentText: res,
		Source:          source,
	}, nil
}
