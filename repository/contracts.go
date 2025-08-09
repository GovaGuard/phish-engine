package repository

import "github.com/holgerson97/phish-engine/entity"

type (
	CampaignRepo interface {
		GetActiveCampaigns() ([]entity.Campaign, error)
		GetCampaigns(string) ([]entity.Campaign, error)
		AddCampaign(entity.Campaign) (entity.Campaign, error)
		UpdateCampaign(entity.Campaign) (entity.Campaign, error)
		DeleteCampaign(string) error
		DeleteAllCampaigns() error
	}
	TargetsRepo interface {
		GetTargets(string) ([]entity.Target, error)
		AddTargets([]entity.Target) ([]entity.Target, error)
		DeleteTarget(string) error
		ChangeTargetState(string, entity.TargetState) error
	}
)
