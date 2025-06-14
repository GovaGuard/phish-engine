package router

import (
	"encoding/json"
	"log"
	"net/http"
)

func (router *Router) GetAllCampaigns(w http.ResponseWriter, r *http.Request) {
	response, err := router.usecase.GetActiveCampaigns()
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

func (router *Router) DeleteAllCampaigns(w http.ResponseWriter, r *http.Request) {
	err := router.usecase.DeleteAllCampaigns()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		log.Println(err)
		return
	}
}
