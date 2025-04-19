package campaigns

import (
	"github.com/google/uuid"
	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/repository/rethinkdb"
)

func NewCampagin(c entity.Campaign) error {
	client, err := rethinkdb.NewClient("localhost:28015")
	if err != nil {
		return err
	}

	// Create the ID for the campaign
	c.ID = uuid.New().String()

	if err := client.AddCampaign(c); err != nil {
		return err
	}

	return nil
}

func GetCampaigns(orgID string) ([]entity.Campaign, error) {
	client, err := rethinkdb.NewClient("localhost:28015")
	if err != nil {
		return nil, err
	}

	campaigns, err := client.GetCampaigns(orgID)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}
