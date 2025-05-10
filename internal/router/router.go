package router

import (
	"net/http"

	"github.com/holgerson97/phish-engine/internal/usecase"
	"github.com/rs/cors"
)

type Router struct {
	usecase *usecase.Usecase
}

func New(c *usecase.Usecase) Router {
	return Router{c}
}

func (router *Router) Run() error {
	mux := http.NewServeMux()

	mux.Handle("GET /api/campaigns", http.HandlerFunc(router.GetCampaigns))
	mux.Handle("POST /api/campaigns", http.HandlerFunc(router.AddCampaign))
	mux.Handle("DELETE /api/campaigns/{id}", http.HandlerFunc(router.DeleteCampaign))
	mux.Handle("GET /api/targets", http.HandlerFunc(router.GetTargets))
	mux.Handle("POST /api/targets", http.HandlerFunc(router.AddTargets))

	handler := cors.AllowAll().Handler(mux)
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		return err
	}

	return nil
}
