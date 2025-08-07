package usecase_test

import (
	"html/template"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/internal/mail"
	"github.com/holgerson97/phish-engine/internal/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func campaignUsecase(t *testing.T) (*usecase.Usecase, *MockCampaignRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockCampaignRepo(mockCtl)
	targetRepo := NewMockTargetsRepo(mockCtl)

	usc := usecase.New(repo, targetRepo, mail.Sender{
		Sender:     "mock@phish-engine.com",
		User:       "mock",
		Password:   "mock",
		Host:       "127.0.0.1",
		SMTPServer: "127.0.0.1:2525",
	})

	return usc, repo
}

func Test_AddCampaign(t *testing.T) {
	usc, repo := campaignUsecase(t)
	c := entity.Campaign{}

	tests := []test{{
		name: "Simple",
		mock: func() {
			repo.EXPECT().AddCampaign(c).Return(c, nil)
		},
		res: c,
		err: nil,
	}}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			v.mock()
			_, err := usc.AddCampaign(c)

			require.ErrorIs(t, err, v.err)
		})
	}
}

func Test_GetCampaigns(t *testing.T) {
	usc, repo := campaignUsecase(t)
	c := []entity.Campaign{{}}

	tests := []test{{
		name: "Simple",
		mock: func() {
			repo.EXPECT().GetCampaigns("1").Return(c, nil)
		},
		res: c,
		err: nil,
	}}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			v.mock()
			campaigns, err := usc.GetCampaigns("1")

			require.Equal(t, c, campaigns)
			require.ErrorIs(t, err, v.err)
		})
	}
}

func Test_GetActiveCampaigns(t *testing.T) {
	usc, repo := campaignUsecase(t)
	c := []entity.Campaign{{}}

	tests := []test{{
		name: "Simple",
		mock: func() {
			repo.EXPECT().GetActiveCampaigns().Return(c, nil)
		},
		res: c,
		err: nil,
	}}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			v.mock()
			campaigns, err := usc.GetActiveCampaigns()

			require.Equal(t, c, campaigns)
			require.ErrorIs(t, err, v.err)
		})
	}
}

func TestUsecase_WorkCampaigns(t *testing.T) {
	usc, repo := campaignUsecase(t)
	tmpl, err := template.New("simpel").Parse("COLL")
	if err != nil {
		t.Fatal("template cannot be parsed")
	}

	empty := []entity.Campaign{}
	simple := []entity.Campaign{{
		ID:             uuid.NewString(),
		CreatorID:      uuid.NewString(),
		OrganizationID: uuid.NewString(),
		Title:          "Simple",
		Status:         entity.CampaignPlanned,
		StartDate:      time.Now(),
		EndDate:        time.Now().Add(time.Second * 15),
		SuccessRate:    0,
		Targets: []entity.Target{{
			ID:             uuid.NewString(),
			OrganizationID: uuid.NewString(),
			EMail:          "mail@phish-engine.com",
			Firstname:      "Gova",
			Surname:        "Guard",
			State:          0,
		}},
		Attack: entity.AttackType{
			ID:     uuid.NewString(),
			Params: map[string]any{},
			Body:   tmpl,
		},
		AttackParams: map[string]any{
			"sender":  "phish@phish-engine.com",
			"subject": "Scam",
		},
	}}

	simpleRunning := simple
	simpleRunning[0].Status = entity.CampaignRunning

	tests := []test{{
		name: "Empty",
		mock: func() {
			repo.EXPECT().GetActiveCampaigns().Return(empty, nil)
		},
		err: nil,
	}, {
		name: "Simple",
		mock: func() {
			repo.EXPECT().GetActiveCampaigns().Return(simple, nil)
			repo.EXPECT().UpdateCampaign(simpleRunning[0]).Return(simpleRunning[0], nil)
		},
		err: nil,
	}}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			v.mock()
			err := usc.WorkCampaigns()
			require.ErrorIs(t, err, v.err)
		})
	}
}
