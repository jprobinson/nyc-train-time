package nyctraintime

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jprobinson/gosubway"
)

func (s *service) getNextTrainDialog(ctx context.Context, ft gosubway.FeedType, line, stop, dir string) string {
	return s.getTrainDialog(ctx, ft, "next", line, stop, dir, 0)
}

func (s *service) getFollowingTrainDialog(ctx context.Context, ft gosubway.FeedType, line, stop, dir string) string {
	return s.getTrainDialog(ctx, ft, "following", line, stop, dir, 1)
}

func (s *service) getTrainDialog(ctx context.Context, ft gosubway.FeedType, name, line, stop, dir string, indx int) string {
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
