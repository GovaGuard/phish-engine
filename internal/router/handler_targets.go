package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/holgerson97/phish-engine/entity"
)

func (router *Router) GetTargets(w http.ResponseWriter, r *http.Request) {
	id, err := parseOrganization(*r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
		return
	}

	response, err := router.usecase.GetTargets(id)
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

func (router *Router) AddTargets(w http.ResponseWriter, r *http.Request) {
	id, err := parseOrganization(*r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var target []entity.Target
	if bytes.HasPrefix(bytes.TrimSpace(body), []byte("[")) {
		if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
	} else {
		var t entity.Target
		if err := json.Unmarshal(body, &t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		target = append(target, t)
	}

	result, err := router.usecase.AddTarget(id, target)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write(resp)
	return
}

func (router *Router) DeleteTarget(w http.ResponseWriter, r *http.Request) {
	targetID := r.PathValue("id")

	if err := router.usecase.DeleteTarget(targetID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed deleting target"))
		return

	}

	w.Write([]byte(fmt.Sprintf("delete %s succesfully", targetID)))

	return
}
