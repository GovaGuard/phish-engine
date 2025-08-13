package router

import (
	"log"
	"net/http"
)

func (router *Router) PhishAction(w http.ResponseWriter, r *http.Request) {
	campaignID := r.PathValue("id")
	targetID := r.PathValue("target_id")

	// TODO: Actual return error type and populate HTTP error Code
	// Issue URL: https://github.com/GovaGuard/phish-engine/issues/24
	if err := router.usecase.TargetPhished(campaignID, targetID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	return
}
