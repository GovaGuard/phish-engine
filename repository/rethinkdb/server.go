package rethinkdb

import (
	"github.com/holgerson97/phish-engine/entity"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Client struct {
	*r.Session
}

const (
	campgainTable = "campaigns"
)

func NewClient(url string) (*Client, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		return nil, err
	}

	return &Client{session}, nil
}

func (cl *Client) GetCampaigns(orgID string) ([]entity.Campaign, error) {
	resp, err := r.Table(campgainTable).Filter(r.Row.Field("organization_id").Eq(orgID)).Run(cl)
	if err != nil {
		return nil, err
	}

	var result []entity.Campaign
	if err := resp.All(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (cl *Client) AddCampaign(c entity.Campaign) error {
	_, err := r.Table(campgainTable).Insert(c).Run(cl)
	if err != nil {
		return err
	}

	return nil
}
