package mongodb

import (
	"context"
	"fmt"

	"github.com/holgerson97/phish-engine/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (cl *Client) GetCampaign(campaignID string) (entity.Campaign, error) {
	coll := cl.Client.Database("main").Collection(campgainTable)
	filter := bson.D{{Key: "creator_id", Value: "314920484783891083"}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "targets", Value: 1}})

	var campaign entity.Campaign

	if err := coll.FindOne(context.TODO(), filter, opts).Decode(&campaign); err != nil {
		fmt.Println(err)
		return campaign, fmt.Errorf("parsing campaign to entity: %w", err)
	}

	return campaign, nil
}

func (cl *Client) AddCampaign(c entity.Campaign) (entity.Campaign, error) {
	coll := cl.Client.Database("main").Collection(campgainTable)

	_, err := coll.InsertOne(context.TODO(), c)
	if err != nil {
		return entity.Campaign{}, fmt.Errorf("adding campaign: %w", err)
	}

	return c, nil
}

func (cl *Client) UpdateCampaignTargets(c entity.Campaign) (entity.Campaign, error) {
	coll := cl.Client.Database("main").Collection(campgainTable)

	filter := bson.D{{Key: "creator_id", Value: "314920484783891083"}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "targets", Value: c.Targets},
	}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return entity.Campaign{}, fmt.Errorf("updating campaign: %w", err)
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
