package mongodb

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
	filter := bson.D{
		{Key: "status", Value: bson.D{
			{Key: "$in", Value: bson.A{entity.CampaignPlanned, entity.CampaignRunning}},
		}},
	}

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

func (cl *Client) AddCampaign(c entity.Campaign) (entity.Campaign, error) {
	coll := cl.Client.Database("main").Collection(campgainTable)

	_, err := coll.InsertOne(context.TODO(), c)
	if err != nil {
		return entity.Campaign{}, fmt.Errorf("adding campaign: %w", err)
	}

	return c, nil
}

func (cl *Client) UpdateCampaign(c entity.Campaign) (entity.Campaign, error) {
	coll := cl.Client.Database("main").Collection(campgainTable)
	filter := bson.D{{Key: "_id", Value: c.ID}}

	_, err := coll.ReplaceOne(context.TODO(), filter, c)
	if err != nil {
		return entity.Campaign{}, fmt.Errorf("updating campaign: %w", err)
	}

	return c, nil
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
