package rethinkdb

import (
	"context"
	"fmt"

	"github.com/holgerson97/phish-engine/entity"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	campgainTable = "campaigns"
)

func (cl *Client) GetActiveCampaigns() ([]entity.Campaign, error) {
	coll := cl.Client.Database("main").Collection(campgainTable)
	filter := bson.D{{Key: "status", Value: entity.StateActive}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	result := []entity.Campaign{}
	if err := cursor.All(context.TODO(), &result); err != nil {
		return nil, fmt.Errorf("parsing campaign to entity: %w", err)
	}

	return result, nil
}

func (cl *Client) GetCampaigns(orgID string) ([]entity.Campaign, error) {
	coll := cl.Client.Database("main").Collection(campgainTable)
	filter := bson.D{{Key: "organization_id", Value: orgID}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	result := []entity.Campaign{}
	if err := cursor.All(context.TODO(), &result); err != nil {
		return nil, fmt.Errorf("parsing campaign to entity: %w", err)
	}

	return result, nil
}

func (cl *Client) AddCampaign(c entity.Campaign) error {
	coll := cl.Client.Database("main").Collection(campgainTable)

	_, err := coll.InsertOne(context.TODO(), c)
	if err != nil {
		return fmt.Errorf("adding campaign: %w", err)
	}

	return nil
}

func (cl *Client) DeleteCampaign(id string) error {
	coll := cl.Client.Database("main").Collection(campgainTable)
	filter := bson.D{{Key: "id", Value: id}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (cl *Client) DeleteAllCampaigns() error {
	coll := cl.Client.Database("main").Collection(campgainTable)
	filter := bson.D{{}}

	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
