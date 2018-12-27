package nyctraintime

import (
	"context"

	"cloud.google.com/go/datastore"
)

type myStop struct {
	Line string
	Stop string
	Dir  string
}

type db struct {
	db *datastore.Client
}

func newDB(ctx context.Context, project string) (*db, error) {
	client, err := datastore.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}
	return &db{db: client}, nil
}

func (d *db) getMyStop(ctx context.Context, userID string) (*myStop, error) {
	var my myStop
	key := datastore.NameKey("MyStop", userID, nil)
	err := d.db.Get(ctx, key, &my)
	return &my, err
}

func (d *db) saveMyStop(ctx context.Context, userID, line, stop, dir string) error {
	key := datastore.NameKey("MyStop", userID, nil)
	_, err := d.db.Put(ctx, key, &myStop{
		Line: line,
		Stop: stop,
		Dir:  dir,
	})
	return err
}
