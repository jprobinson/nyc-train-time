package gosubway

import (
	"sort"
	"strings"
	"time"

	_ "github.com/jprobinson/gtfs/nyct_subway_proto"
	"github.com/jprobinson/gtfs/transit_realtime"
)

type FeedMessage struct {
	transit_realtime.FeedMessage
}

type Alert struct {
	transit_realtime.Alert
}

type StopTimeUpdate struct {
	transit_realtime.TripUpdate_StopTimeUpdate
}

type FeedType int

const (
	NumberedFeed FeedType = 1
	BlueFeed     FeedType = 26
	YellowFeed   FeedType = 16
	OrangeFeed   FeedType = 21
	LFeed        FeedType = 2
	GFeed        FeedType = 31
	SevenFeed    FeedType = 51
	BrownFeed    FeedType = 36
)

// Trains will accept a stopId plus a train line (found here: http://web.mta.info/developers/data/nyct/subway/google_transit.zip)
// and returns a list of updates from northbound and southbound trains
func (f *FeedMessage) Trains(stopId, line string) (alerts []*Alert, northbound, southbound []*StopTimeUpdate) {
	for _, ent := range f.Entity {
		if ent.TripUpdate != nil {
			for _, upd := range ent.TripUpdate.StopTimeUpdate {
				if strings.HasPrefix(*upd.StopId, stopId) &&
					line == *ent.TripUpdate.Trip.RouteId {
					if strings.HasSuffix(*upd.StopId, "N") {
						northbound = append(northbound, &StopTimeUpdate{*upd})
					} else if strings.HasSuffix(*upd.StopId, "S") {
						southbound = append(southbound, &StopTimeUpdate{*upd})
					}
				}
			}
			if ent.Alert != nil {
				alerts = append(alerts, &Alert{*ent.Alert})
			}
		}

	}
	return alerts, northbound, southbound
}

// NextTrainTimes will return an ordered slice of upcoming train departure times
// in either direction.
func (f *FeedMessage) NextTrainTimes(stopId, line string) (alerts []*Alert, northbound, southbound []time.Time) {
	alerts, north, south := f.Trains(stopId, line)
	northbound = NextTrainTimes(north)
	southbound = NextTrainTimes(south)
	return alerts, northbound, southbound
}

// NextTrainTimes will extract the departure times from the given
// update slice, order and return them.
func NextTrainTimes(updates []*StopTimeUpdate) []time.Time {
	var times []time.Time

	for _, upd := range updates {
		if upd.Departure == nil {
			continue
		}
		unix := *upd.Departure.Time
		if upd.Departure.Delay != nil {
			unix += int64(*upd.Departure.Delay * 1000)
		}
		dept := time.Unix(unix, 0)
		if dept.After(time.Now()) {
			times = append(times, dept)
		}
	}
	sort.Sort(timeSlice(times))
	if len(times) > 5 {
		times = times[:5]
	}
	return times
}

type timeSlice []time.Time

func (t timeSlice) Len() int {
	return len(t)
}

func (t timeSlice) Less(i, j int) bool {
	return t[i].Before(t[j])
}

func (t timeSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
