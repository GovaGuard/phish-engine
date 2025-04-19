package router

import (
	"encoding/json"
	"net/http"

	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/internal/usecase/campaigns"
)

func GetCampaigns(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")

	response, err := campaigns.GetCampaigns(orgID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		return
	}

	body, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed creating payload"))
		return
	}

	w.Write(body)
}

func AddCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign entity.Campaign
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		return
	}

	if err := campaigns.NewCampagin(campaign); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed creating campaign"))
		return

	}

	return
}

func Run() error {
	mux := http.NewServeMux()

	gc := http.HandlerFunc(GetCampaigns)
	ac := http.HandlerFunc(AddCampaign)
	mux.Handle("/api/getcampaigns", gc)
	mux.Handle("/api/addcampaigns", ac)

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		return err
	}

	return nil
}
