package router

import (
	"encoding/json"
	"net/http"

	"github.com/holgerson97/phish-engine/entity"
)

func GetCampaigns(w http.ResponseWriter, r *http.Request) {
	mock := entity.Campaign{Name: "TestingCampaign", State: "Active", OwnerID: "123"}
	resp, err := json.Marshal(mock)
	if err != nil {
		return
	}

	w.Write(resp)
}

func Run() error {
	mux := http.NewServeMux()

	gc := http.HandlerFunc(GetCampaigns)
	mux.Handle("/api/getcampaigns", gc)

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		return err
	}

	return nil
}
