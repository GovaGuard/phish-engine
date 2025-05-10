package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/holgerson97/phish-engine/entity"
)

func (usc *Usecase) AddTarget(org string, t []entity.Target) ([]entity.Target, error) {
	// Create the ID for the target
	for key := range t {
		t[key].ID = uuid.New().String()
		t[key].OrganizationID = org
	}

	fmt.Println(t)
	result, err := usc.repository.AddTarget(t)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (usc *Usecase) GetTargets(orgID string) ([]entity.Target, error) {
	targets, err := usc.repository.GetTargets(orgID)
	if err != nil {
		return nil, err
	}

	return targets, nil
}

func (usc *Usecase) DeleteTarget(id string) error {
	err := usc.repository.DeleteTarget(id)
	if err != nil {
		return err
	}

	return nil
}
