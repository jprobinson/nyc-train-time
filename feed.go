package nyctraintime

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NYTimes/gizmo/server/kit"

	"github.com/jprobinson/gtfs/mta"
	"github.com/jprobinson/gtfs/transit_realtime"
)

const (
	timeout     = 1 * time.Second
	maxAttempts = 10
	backoffStep = 50 * time.Millisecond
)

// retries until it hits max attempts or a context timeout
func getFeed(ctx context.Context, hc *http.Client, key string, ft mta.FeedType) (*transit_realtime.FeedMessage, error) {
	var (
		feed *transit_realtime.FeedMessage
		err  error
	)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		// retry backoff
		time.Sleep(time.Duration((attempt - 1)) * backoffStep)
		// attempt to get feed
		feed, err = mta.GetNYCSubwayFeed(ctx, hc, key, ft)
		if err == nil ||
			(err != nil && strings.Contains(err.Error(), "deadline exceeded")) {
			break
		}
		kit.LogErrorMsg(ctx, err, fmt.Sprintf("unable to get mta feed on attempt %d", attempt))
	}
	return feed, err
}

func parseFeed(line string) (mta.FeedType, error) {
	var ft mta.FeedType
	switch line {
	case "1", "2", "3", "4", "5", "6":
		ft = mta.NumberedFeed
	case "N", "Q", "R", "W":
		ft = mta.YellowFeed
	case "B", "D", "F", "M":
		ft = mta.OrangeFeed
	case "A", "C", "E":
		ft = mta.BlueFeed
	case "J", "Z":
		ft = mta.BrownFeed
	case "L":
		ft = mta.LFeed
	case "7":
		ft = mta.SevenFeed
	case "G":
		ft = mta.GFeed
	default:
		return mta.LFeed, errBadRequest
	}
	return ft, nil
}

var errBadRequest = kit.NewJSONStatusResponse(map[string]string{
	"error": "bad request"}, http.StatusBadRequest)
