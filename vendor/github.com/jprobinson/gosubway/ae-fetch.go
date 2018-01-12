// +build appengine

package gosubway

import (
	"io/ioutil"
	"strconv"

	"github.com/golang/protobuf/proto"
	_ "github.com/jprobinson/gtfs/nyct_subway_proto"
	"github.com/jprobinson/gtfs/transit_realtime"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

// GetFeed takes an API key generated from http://datamine.mta.info/user/register
// and a boolean specifying which feed (1,2,3,4,5,6,S trains OR L train) and
// it will return a transit_realtime.FeedMessage with NYCT extensions.
func GetFeed(ctx context.Context, key string, ft FeedType) (*FeedMessage, error) {
	url := "http://datamine.mta.info/mta_esi.php?key=" + key +
		"&feed_id=" + strconv.Itoa(int(ft))
	cl := urlfetch.Client(ctx)
	resp, err := cl.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	transit := &transit_realtime.FeedMessage{}
	err = proto.Unmarshal(body, transit)
	if err != nil {
		return nil, err
	}
	return &FeedMessage{*transit}, nil
}
