package usecase

import "github.com/holgerson97/phish-engine/entity"

type (
	Campaigns interface {
		AddCampaign(entity.Campaign) error
		GetCampaigns(string) ([]entity.Campaign, error)
		GetActiveCampaigns() ([]entity.Campaign, error)
		DeleteCampaign(string) error
		DeleteAllCampaigns() error
		WorkCampaign() error
	}
	Targets interface {
		AddTarget(string, []entity.Target) ([]entity.Target, error)
		GetTargets(string) ([]entity.Target, error)
		DeleteTarget(string) error
	}
)
