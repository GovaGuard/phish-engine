package usecase

import (
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

func (usc *Usecase) AddCampaign(c entity.Campaign) error {
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

func (usc *Usecase) WorkCampaigns() ([]entity.Campaign, error) {
	campaigns, err := usc.repository.GetActiveCampaigns()
	if err != nil {
		return nil, err
	}

	noErrorCampaigns := []entity.Campaign{}
	for values := range slices.Values(campaigns) {
		for target := range slices.Values(values.Targets) {
			if target.State != entity.StateCompleted {

				attack, err := entity.NewAttackDownload(values.Attack)
				if err != nil {
					return noErrorCampaigns, err
				}

				body, err := attack.Template(values.AttackParams)
				if err != nil {
					return noErrorCampaigns, err
				}

				// In order to be able to send out multi targets in the
				// future this function already takes in mulitple
				// recipients
				m := attack.GenerateMail([]string{target.EMail}, body)

				state := target.State
				if err := usc.smtp.SendMail(m); err != nil {
					log.Println(err)
					state = entity.StateError
					continue
				}

				state = entity.StateCompleted
				if err := usc.targetRepository.ChangeTargetState(target.ID, state); err != nil {
					log.Println(err)
				}
			}
		}

		noErrorCampaigns = append(noErrorCampaigns, values)
	}

	return noErrorCampaigns, nil
}
