package router

import (
	"encoding/json"
	"net/http"

	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/internal/usecase/campaigns"
)

type Router struct {
	usecase *campaigns.Usecase
}

func New(c *campaigns.Usecase) Router {
	return Router{c}
}

func (router *Router) GetCampaigns(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")

	response, err := router.usecase.GetCampaigns(orgID)
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

func (router *Router) AddCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign entity.Campaign
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		return
	}

	if err := router.usecase.AddCampaign(campaign); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed creating campaign"))
		return

	}

	return
}

func (router *Router) Run() error {
	mux := http.NewServeMux()

	mux.Handle("GET /api/campaigns", http.HandlerFunc(router.GetCampaigns))
	mux.Handle("POST /api/campaigns", http.HandlerFunc(router.AddCampaign))

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		return err
	}

	return nil
}
