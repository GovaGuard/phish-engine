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
	id, err := parseOrganization(*r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
		return
	}

	userID, err := parseToken(*r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
		return
	}

	var campaign entity.Campaign
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		log.Println(err)
		return
	}

	// TODO: This maybe needs to be reworked
	campaign.OrganizationID = id
	campaign.CreatorID = userID

	result, err := router.usecase.AddCampaign(campaign)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed creating campaign"))
		log.Println(err)
		return
	}

	resp, err := json.Marshal(result)

	w.Write(resp)

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
