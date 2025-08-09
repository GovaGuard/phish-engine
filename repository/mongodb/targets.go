package mongodb

import (
	"context"
	"fmt"

	"github.com/holgerson97/phish-engine/entity"
	"go.mongodb.org/mongo-driver/bson"
)

const targetTable = "targets"

func (cl *Client) GetTargets(orgID string) ([]entity.Target, error) {
	coll := cl.Client.Database("main").Collection(targetTable)
	filter := bson.D{{Key: "organization_id", Value: orgID}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	result := []entity.Target{}
	if err := cursor.All(context.TODO(), &result); err != nil {
		return nil, fmt.Errorf("parsing target to entity: %w", err)
	}

	return result, nil
}

func (cl *Client) AddTargets(t []entity.Target) ([]entity.Target, error) {
	coll := cl.Client.Database("main").Collection(targetTable)

	if len(t) == 1 {
		_, err := coll.InsertOne(context.TODO(), t[0])
		if err != nil {
			return []entity.Target{}, fmt.Errorf("adding target: %w", err)
		}

		return t, nil
	}

	docs := []any{t}
	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return []entity.Target{}, fmt.Errorf("adding targets: %w", err)
	}

	return t, nil
}

func (cl *Client) DeleteTarget(id string) error {
	coll := cl.Client.Database("main").Collection(targetTable)
	filter := bson.D{{Key: "id", Value: id}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (cl *Client) ChangeTargetState(id string, state entity.TargetState) error {
	coll := cl.Client.Database("main").Collection(targetTable)
	filter := bson.D{{Key: "id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: state}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
