package nyctraintime

import (
	"context"

	"google.golang.org/appengine/datastore"
)

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
