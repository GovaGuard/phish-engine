package rethinkdb

import (
	"github.com/holgerson97/phish-engine/entity"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

const targetTable = "targets"

func (cl *Client) GetTargets(orgID string) ([]entity.Target, error) {
	resp, err := r.Table(targetTable).Filter(r.Row.Field("organization_id").Eq(orgID)).Run(cl)
	if err != nil {
		return nil, err
	}

	result := []entity.Target{}
	if err := resp.All(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (cl *Client) AddTarget(t []entity.Target) ([]entity.Target, error) {
	resp, err := r.Table(targetTable).Insert(t).Run(cl)
	if err != nil {
		return nil, err
	}

	result := []entity.Target{}
	if err := resp.All(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (cl *Client) DeleteTarget(id string) error {
	_, err := r.Table(targetTable).Filter(r.Row.Field("id").Eq(id)).Delete().Run(cl)
	if err != nil {
		return err
	}

	return nil
}
