package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	"github.com/holgerson97/phish-engine/internal/mail"
)

type Campaign struct {
	ID             string         `binding:"required" json:"id"`
	CreatorID      string         `binding:"required" json:"creator_id"`
	OrganizationID string         `binding:"required" json:"organization_id"`
	Title          string         `binding:"required" json:"title"`
	Status         string         `binding:"required" json:"status"`
	StartDate      time.Time      `binding:"required" json:"start_date"`
	EndDate        time.Time      `binding:"required" json:"end_date"`
	SuccessRate    int16          `binding:"required" json:"success_rate"`
	Targets        []Target       `binding:"required" json:"targets"`
	Attack         string         `binding:"required" json:"attack"`
	AttackParams   map[string]any `binding:"required" json:"attack_params"`
}

type Attack interface {
	Template(map[string]any) (string, error)
	GenerateMail([]string, string) mail.Mail
}

type AttackDownload struct {
	Sender       string `binding:"required" json:"sender"`
	Subject      string `binding:"required" json:"subject"`
	BodyTemplate *template.Template
}

func NewAttackDownload(body string) (AttackDownload, error) {
	var attackBody map[string]any
	var attack AttackDownload

	if err := json.Unmarshal([]byte(body), &attackBody); err != nil {
		return attack, err
	}

	t, ok := attackBody["body"].(string)
	if !ok {
		return AttackDownload{}, fmt.Errorf("failed unmarshalling attack body")
	}

	tmpl, err := template.New("phishingEmail").Parse(t)
	if err != nil {
		return AttackDownload{}, err
	}

	return AttackDownload{BodyTemplate: tmpl}, nil
}

func (attack *AttackDownload) Template(params map[string]any) (string, error) {
	buf := &bytes.Buffer{}

	if err := attack.BodyTemplate.Execute(buf, params); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (attack *AttackDownload) GenerateMail(t []string, body string) mail.Mail {
	m := mail.NewPlainMail(attack.Sender, attack.Subject, body, t)

	return m
}
