package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/holgerson97/phish-engine/entity"
)

func (router *Router) GetCampaigns(w http.ResponseWriter, r *http.Request) {
	id, err := parseOrganization(*r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
		return
	}

	response, err := router.usecase.GetCampaigns(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		log.Println(err)
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
		log.Println(err)
		return
	}

	if err := router.usecase.AddCampaign(campaign); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed creating campaign"))
		return

	}

	return
}

func (router *Router) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	campaignID := r.PathValue("id")

	err := router.usecase.DeleteCampaign(campaignID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed deleting campaign"))
		return
	}

	return
}
