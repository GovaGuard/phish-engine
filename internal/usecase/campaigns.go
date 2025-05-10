package usecase

import (
	"github.com/google/uuid"
	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/repository/rethinkdb"
)

type Usecase struct {
	repository *rethinkdb.Client
}

func New(url string) (*Usecase, error) {
	client, err := rethinkdb.NewClient(url)
	if err != nil {
		return nil, err
	}

	return &Usecase{repository: client}, nil
}

func (usc *Usecase) AddCampaign(c entity.Campaign) error {
	// Create the ID for the campaign
	c.ID = uuid.New().String()

	if err := usc.repository.AddCampaign(c); err != nil {
		return err
	}

	return nil
}

func (usc *Usecase) GetCampaigns(orgID string) ([]entity.Campaign, error) {
	campaigns, err := usc.repository.GetCampaigns(orgID)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (usc *Usecase) DeleteCampaign(id string) error {
	if err := usc.DeleteCampaign(id); err != nil {
		return err
	}

	return nil
}
