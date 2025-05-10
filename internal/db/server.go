package db

import (
	"slices"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func New(url string) (*r.Session, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		return nil, err
	}

	c, err := r.TableList().Run(session)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	var tables []string
	if err := c.All(&tables); err != nil {
		return nil, err
	}

	if !slices.Contains(tables, "campaigns") {
		_, err = r.TableCreate("campaigns").Run(session)
		if err != nil {
			return nil, err
		}
	}

	return session, nil
}
