package router

import (
	"fmt"
	"net/http"
)

func (router *Router) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	campaignID := r.PathValue("id")

	if err := router.usecase.DeleteCampaign(campaignID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed creating campaign"))
		return

	}

	w.Write([]byte(fmt.Sprintf("delete %s succesfully", campaignID)))

	return
}
