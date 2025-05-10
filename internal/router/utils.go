package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: Make function type safety
// Issue URL: https://github.com/GovaGuard/phish-engine/issues/5
func parseOrganization(r http.Request) (string, error) {
	token := r.Header.Get("Authorization")

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

	id, ok := m["urn:zitadel:iam:user:resourceowner:id"].(string)
	if !ok {
		return "", fmt.Errorf("Organization doesn't exist")
	}

	return id, nil
}
