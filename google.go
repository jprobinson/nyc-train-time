package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/NYTimes/marvin"
	"github.com/jprobinson/dialogflow"
	"github.com/jprobinson/gosubway"
)

func NewGoogleServer() dialogflow.GoogleActionServer {
	return &google{key: os.Getenv("MTA_KEY")}
}

type google struct {
	key string
}

func (g *google) Actions() map[string]dialogflow.GoogleActionHandler {
	return map[string]dialogflow.GoogleActionHandler{
		"my_next_train_request":      g.myTrain,
		"my_following_train_request": g.myFollowingTrain,
		"save_my_stop_request":       g.saveMyStopAction,
		"next_train_request":         g.nextTrain,
		"following_train_request":    g.followingTrain,
	}
}

func (g *google) myTrain(ctx context.Context, r *dailogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
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
		s.getNextTrainDialog(ctx, ft, mys.Line, mys.Stop, mys.Dir))
}

func (g *google) myFollowingTrain(ctx context.Context, r *dailogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
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
		s.getFollowingTrainDialog(ctx, ft, mys.Line, mys.Stop, mys.Dir))
}

func (g *google) saveMyStopAction(ctx context.Context, r *dailogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	uid := r.OriginalRequest.Data.User.UserID
	if uid == "" {
		return simpleGoogleResponse("Sorry, you need to be logged in for that to work")
	}

	line := strings.ToUpper(r.Result.Parameters["subway-line"].(string))
	stop := r.Result.Parameters["subway-stop"].(string)
	dir := r.Result.Parameters["subway-direction"].(string)

	err = saveMyStop(ctx, uid, line, stop, dir)
	if err != nil {
		return nil, marvin.NewJSONStatusResponse(map[string]string{
			"error": "unable to complete request: " + err.Error(),
		}, http.StatusInternalServerError)
	}
	return simpleGoogleResponse(fmt.Sprintf(
		"Successfully saved your stop, %s bound %s trains at %s. To update your stop again, ask NYC Train Time to \"save my stop\". ",
		dir, line, stop))
}

func (g *google) nextTrain(ctx context.Context, r *dailogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	line := strings.ToUpper(r.Result.Parameters["subway-line"].(string))
	stop := r.Result.Parameters["subway-stop"].(string)
	dir := r.Result.Parameters["subway-direction"].(string)

	ft, err := parseFeed(line)
	if err != nil {
		log.Debugf(ctx, "unable to parse line: %s", line)
		res = fmt.Sprintf("sorry, the %s line is not available yet", line)
		break
	}

	res = s.getNextTrainDialog(ctx, ft, line, stop, dir) +
		" If you would like me to remember your stop, ask NYC Train Time to \"save my stop\" and then ask for MY stop next time. "
}

func (g *google) followingTrain(ctx context.Context, r *dailogflow.GoogleRequest) (*dialogflow.GoogleFulfillmentResponse, error) {
	line := strings.ToUpper(r.Result.Parameters["subway-line"].(string))
	stop := r.Result.Parameters["subway-stop"].(string)
	dir := r.Result.Parameters["subway-direction"].(string)
	ft, err := parseFeed(line)
	if err != nil {
		log.Debugf(ctx, "unable to parse line: %s", line)
		res = fmt.Sprintf("Sorry, the %s line is not available yet.", line)
		break
	}
	res = s.getFollowingTrainDialog(ctx, ft, line, stop, dir)
}

func (s *google) getNextTrainDialog(ctx context.Context, ft gosubway.FeedType, line, stop, dir string) string {
	return s.getTrainDialog(ctx, ft, "next", line, stop, dir, 0)
}

func (s *google) getFollowingTrainDialog(ctx context.Context, ft gosubway.FeedType, line, stop, dir string) string {
	return s.getTrainDialog(ctx, ft, "following", line, stop, dir, 1)
}

func (s *google) getTrainDialog(ctx context.Context, ft gosubway.FeedType, name, line, stop, dir string, indx int) string {
	feed, err := getFeed(ctx, s.key, ft)
	if err != nil {
		return fmt.Sprintf("Sorry, I'm having problems getting the subway feed. ")
	}

	stopLine, ok := stopNameToID[stop]
	if !ok {
		return fmt.Sprintf("Sorry, I didn't recognise the stop \"%s\". ", stop)
	}

	stopID, ok := stopLine[line]
	if !ok {
		return fmt.Sprintf("Sorry, I didn't recognise \"%s\" as a part of the %s line. ",
			stop, line)
	}

	_, north, south := feed.NextTrainTimes(stopID, line)

	var trains []time.Time
	if trainDirs[line]["northbound"] == dir || dir == "uptown" || dir == "Northbound" {
		trains = north
	} else {
		trains = south
	}

	if len(trains) < indx+1 {
		return fmt.Sprintf("Sorry, the %s train time is not available for %s bound %s trains at %s. ",
			name, dir, line, stop)
	}

	out := timeSpeak(trains[indx], name, line, stop, dir)
	if len(trains) >= indx+2 {
		out += timeSpeak(trains[indx+1], "following", line, stop, dir)
	}
	return out
}

func timeSpeak(t time.Time, name, line, stop, dir string) string {
	diff := t.Sub(time.Now().UTC())
	mins := strconv.Itoa(int(diff.Minutes()))
	secs := strconv.Itoa(int(diff.Seconds()) % 60)
	out := fmt.Sprintf("The %s %s train will leave %s towards %s in ",
		name, line, stop, dir)
	if mins != "0" {
		out += mins + " minutes and "
	}
	out += secs + " seconds. "
	return out
}

const source = "Where's The Train (NYC)"

func simpleGoogleResponse(res string) (*dialogflow.GoogleFulfillmentResponse, error) {
	return &dialogflow.GoogleFulfillmentResponse{
		Speech:      res,
		DisplayText: res,
		Source:      source,
	}, nil
}

type myStop struct {
	Line string
	Stop string
	Dir  string
}

func getMyStop(ctx context.Context, userID string) (*myStop, error) {
	var my myStop
	err := datastore.Get(ctx, datastore.NewKey(ctx, "MyStop", userID, 0, nil), &my)
	return &my, err
}

func saveMyStop(ctx context.Context, userID, line, stop, dir string) error {
	_, err := datastore.Put(ctx, datastore.NewKey(ctx, "MyStop", userID, 0, nil), &myStop{
		Line: line,
		Stop: stop,
		Dir:  dir,
	})
	return err
}
