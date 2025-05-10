package router

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/holgerson97/phish-engine/entity"
	"github.com/holgerson97/phish-engine/internal/usecase/campaigns"
	"github.com/rs/cors"
)

type Router struct {
	usecase *campaigns.Usecase
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Request-Method", "*")
}

func New(c *campaigns.Usecase) Router {
	return Router{c}
}

// TODO: Make function type safety
func ParseOrganization(token string) (string, error) {
	url := "https://development-8jjgdu.us1.zitadel.cloud/oidc/v1/userinfo"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	m := make(map[string]any)
	if err := json.Unmarshal(body, &m); err != nil {
		return "", err
	}

	return m["urn:zitadel:iam:user:resourceowner:id"].(string), nil
}

func (router *Router) GetCampaigns(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Authorization")
	id, err := ParseOrganization(accessToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Token"))
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

func (router *Router) Run() error {
	mux := http.NewServeMux()

	mux.Handle("GET /api/campaigns", http.HandlerFunc(router.GetCampaigns))
	mux.Handle("POST /api/campaigns", http.HandlerFunc(router.AddCampaign))
	mux.Handle("DELETE /api/campaigns/{id}", http.HandlerFunc(router.DeleteCampaign))

	handler := cors.AllowAll().Handler(mux)
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		return err
	}

	return nil
}
