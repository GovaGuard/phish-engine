package usecase

import (
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/internal/mail"
	"github.com/holgerson97/phish-engine/repository"
)

// TODO: Add error for campaign is not runngin anymore
// Issue URL: https://github.com/GovaGuard/phish-engine/issues/26
type CampaignNotFoundError struct {
	ID string
}

func (e *CampaignNotFoundError) Error() string {
	return fmt.Sprintf("campaign %s not found", e.ID)
}

type TargetNotFoundInCampaign struct {
	CampaignID, ID string
}

func (e *TargetNotFoundInCampaign) Error() string {
	return fmt.Sprintf("campaign %s does not contain target %s", e.CampaignID, e.ID)
}

type Usecase struct {
	smtp             mail.Sender
	repository       repository.CampaignRepo
	targetRepository repository.TargetsRepo
}

func New(c repository.CampaignRepo, t repository.TargetsRepo, m mail.Sender) *Usecase {
	return &Usecase{repository: c, targetRepository: t, smtp: m}
}

func (usc *Usecase) AddCampaign(c entity.Campaign) (entity.Campaign, error) {
	c.ID = uuid.New().String()
	c.Status = entity.CampaignPlanned

	// TODO: Right now we only support invoice phishing with fixed params, these need
	// to be refactored
	c.AttackParams = map[string]any{
		"sender":  "phish@phish-engine.com",
		"subject": "Invoice Payment",
	}

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

func (usc *Usecase) TargetPhished(campaignID string, targetID string) error {
	campaign, err := usc.repository.GetCampaign(campaignID)
	if err != nil {
		return &CampaignNotFoundError{ID: campaignID}
	}

	for k, t := range campaign.Targets {
		if t.ID == targetID {

			campaign.Targets[k].State = entity.StateSuccess
			campaign, err = usc.repository.UpdateCampaignTargets(campaign)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return &TargetNotFoundInCampaign{CampaignID: campaignID, ID: targetID}
}

func (usc *Usecase) WorkCampaigns() error {
	campaigns, err := usc.repository.GetActiveCampaigns()
	if err != nil {
		return err
	}

	for campaign := range slices.Values(campaigns) {

		// TODO: Move this to a DB querry
		campaign = usc.CheckState(campaign)
		if campaign.Status == entity.CampaignRunning {
			usc.RunCampaign(campaign)
		}

		_, err = usc.repository.UpdateCampaignStatus(campaign)
		if err != nil {
			return fmt.Errorf("failed updating campaign: %s %w", campaign.ID, err)
		}

	}
	return nil
}

func (usc *Usecase) RunCampaign(c entity.Campaign) error {
	errorFunc := func(c entity.Campaign, err error) {
		log.Println(fmt.Errorf("error occured in campaign: %s: %s", c.ID, err))
		c.Status = entity.CampaignError
		_, err = usc.repository.UpdateCampaignStatus(c)
		if err != nil {
			log.Println(err)
		}
	}

	// TODO: Filter out Targets that have been already hit
	// Maybe this can be done cleaner, it also should be a usecase function
	openTargets := []entity.Target{}
	for t := range slices.Values(c.Targets) {
		if t.State == entity.StateActive {
			openTargets = append(openTargets, t)
		}
	}

	attack, err := c.Attack.GenerateMail(c.AttackParams, openTargets)
	if err != nil {
		errorFunc(c, err)
	}

	if err := usc.smtp.SendMail(attack); err != nil {
		errorFunc(c, err)
	}

	for key := range slices.All(openTargets) {
		openTargets[key].State = entity.StateCompleted
	}

	return nil
}

func (usc *Usecase) EvaluateSuccessRate(c entity.Campaign) entity.Campaign {
	success := 0
	for t := range slices.Values(c.Targets) {
		if t.State == entity.StateSuccess {
			success++
		}
	}

	succesRate := (success / len(c.Targets)) * 100
	c.SuccessRate = int16(succesRate)

	return c
}

func (usc *Usecase) CheckState(c entity.Campaign) entity.Campaign {
	if c.StartDate.Before(time.Now()) {
		c.Status = entity.CampaignRunning
	}

	if c.StartDate.After(time.Now()) {
		c.Status = entity.CampaignPlanned
	}

	if c.EndDate.Before(time.Now()) {
		c.Status = entity.CampaignCompleted
	}

	return c
}
