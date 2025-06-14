package rethinkdb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Client
}

func NewClient(ctx context.Context, uri string) (*Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

func (c *Client) Setup() error {
	db := c.Client.Database("main")

	if err := db.CreateCollection(context.TODO(), campgainTable); err != nil {
		return err
	}

	if err := db.CreateCollection(context.TODO(), targetTable); err != nil {
		return err
	}

	return nil
}
