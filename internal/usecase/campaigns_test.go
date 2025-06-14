package usecase_test

import (
	"testing"

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

	usc := usecase.New(repo, targetRepo, mail.Sender{})

	return usc, repo
}

func Test_AddCampaign(t *testing.T) {
	usc, repo := campaignUsecase(t)
	c := entity.Campaign{}

	tests := []test{{
		name: "Simple",
		mock: func() {
			repo.EXPECT().AddCampaign(c).Return(nil)
		},
		res: c,
		err: nil,
	}}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			v.mock()
			err := usc.AddCampaign(c)

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
			campaigns, err := usc.WorkCampaigns()

			require.Equal(t, c, campaigns)
			require.ErrorIs(t, err, v.err)
		})
	}
}

// func TestUsecase_WorkCampaigns(t *testing.T) {
// 	r, err := rethinkdb.NewClient(context.TODO(), "localhost:28015")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	start, _ := time.Parse(time.RFC3339, time.Now().String())
//
// 	attackMap := map[string]any{
// 		"sender":  "ceo@phish-engine.com",
// 		"subject": "Unit-Test",
// 		"body":    entity.EmailTemplate,
// 	}
//
// 	attack, err := json.Marshal(attackMap)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
//
// 	r.AddCampaign(entity.Campaign{
// 		ID:             uuid.NewString(),
// 		CreatorID:      uuid.NewString(),
// 		OrganizationID: uuid.NewString(),
// 		Title:          fmt.Sprintf("test-campaign-%s", uuid.NewString()),
// 		Status:         "active",
// 		StartDate:      start,
// 		EndDate:        start,
// 		Attack:         string(attack),
// 		SuccessRate:    1,
// 		AttackParams:   map[string]any{"EmployeeName": "Rainer Winkler", "DownloadLink": "govaguard.com", "AttachmentName": "Invoice.pdf", "CompanyName": "GovaGuard"},
// 		Targets: []entity.Target{{
// 			ID:             uuid.NewString(),
// 			OrganizationID: uuid.NewString(),
// 			EMail:          "target@phish-engine.com",
// 			Firstname:      "Jon",
// 			Surname:        "Doe",
// 			State:          entity.StateActive,
// 		}},
// 	})
//
// 	s := mail.Sender{
// 		Sender:     "mock@phish-engine.com",
// 		User:       "mock",
// 		Password:   "mock",
// 		Host:       "127.0.0.1",
// 		SMTPServer: "127.0.0.1:2525",
// 	}
//
// 	type fields struct {
// 		repository *rethinkdb.Client
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		{name: "Simple", fields: fields{r}, wantErr: false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			usc := &Usecase{
// 				smtp:       s,
// 				repository: tt.fields.repository,
// 			}
// 			if err := usc.WorkCampaigns(); (err != nil) != tt.wantErr {
// 				t.Errorf("Usecase.WorkCampaigns() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
