package usecase

import (
	"fmt"
	"log"
	"slices"

	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/internal/mail"
	"github.com/holgerson97/phish-engine/repository"
)

type Usecase struct {
	smtp             mail.Sender
	repository       repository.CampaignRepo
	targetRepository repository.TargetsRepo
}

func New(c repository.CampaignRepo, t repository.TargetsRepo, m mail.Sender) *Usecase {
	return &Usecase{repository: c, targetRepository: t, smtp: m}
}

func (usc *Usecase) AddCampaign(c entity.Campaign) (entity.Campaign, error) {
	c.Status = entity.CampaignPlanned

	cmp, err := usc.repository.AddCampaign(c)
	if err != nil {
		return cmp, err
	}

	return cmp, nil
}

func (usc *Usecase) GetCampaigns(orgID string) ([]entity.Campaign, error) {
	campaigns, err := usc.repository.GetCampaigns(orgID)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (usc *Usecase) GetActiveCampaigns() ([]entity.Campaign, error) {
	campaigns, err := usc.repository.GetActiveCampaigns()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (usc *Usecase) DeleteCampaign(id string) error {
	return usc.repository.DeleteCampaign(id)
}

func (usc *Usecase) DeleteAllCampaigns() error {
	return usc.repository.DeleteAllCampaigns()
}

func (usc *Usecase) WorkCampaigns() error {
	campaigns, err := usc.repository.GetActiveCampaigns()
	if err != nil {
		return err
	}

	errorFunc := func(c entity.Campaign, err error) {
		log.Println("error occured", err)
		c.Status = entity.CampaignError
		_, err = usc.repository.UpdateCampaign(c)
		if err != nil {
			log.Println(err)
		}
	}

	for values := range slices.Values(campaigns) {

		values.Status = entity.CampaignRunning

		attack, err := values.Attack.GenerateMail(values.AttackParams, values.Targets)
		if err != nil {
			errorFunc(values, err)
			continue
		}

		if err := usc.smtp.SendMail(attack); err != nil {
			errorFunc(values, err)
			continue
		}

		log.Println(
			fmt.Sprintf(
				"campaign %s switches from %s to %s", values.ID, values.Status, entity.CampaignCompleted))

		values.Status = entity.CampaignCompleted
		_, err = usc.repository.UpdateCampaign(values)
		if err != nil {
			return fmt.Errorf("failed updating campaign: %s %w", values.ID, err)
		}

	}

	return nil
}
