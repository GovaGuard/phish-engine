package entity

import (
	"bytes"
	"errors"
	"html/template"
	"maps"
	"reflect"
	"slices"
	"time"

	"github.com/holgerson97/phish-engine/internal/mail"
)

type CampaignState int

const (
	CampaignPlanned CampaignState = iota
	CampaignRunning
	CampaignCompleted
	CampaignPaused
	CampaignError
	CampaignArchived
	CampaignUnknown
)

var campaignStateName = map[CampaignState]string{
	CampaignCompleted: "completed",
	CampaignRunning:   "running",
	CampaignPlanned:   "scheduled",
	CampaignPaused:    "paused",
	CampaignError:     "error",
	CampaignArchived:  "archived",
	CampaignUnknown:   "unknown",
}

func (cs CampaignState) String() string {
	return campaignStateName[cs]
}

type Campaign struct {
	ID             string         `json:"id" bson:"_id,omitempty"`
	CreatorID      string         `json:"creator_id" bson:"creator_id"`
	OrganizationID string         `json:"organization_id" bson:"organization_id"`
	Title          string         `binding:"required" json:"title" bson:"title"`
	Status         CampaignState  `json:"status" bson:"status"`
	StartDate      time.Time      `binding:"required" json:"start_date" bson:"start_date"`
	EndDate        time.Time      `binding:"required" json:"end_date" bson:"end_date"`
	SuccessRate    int16          `json:"success_rate" bson:"success_rate"`
	Targets        []Target       `json:"targets" bson:"targets"`
	Attack         AttackType     `json:"attack" bson:"attack"`
	AttackParams   map[string]any `json:"attack_params" bson:"attack_params"`
}

type AttackType struct {
	ID     string             `json:"id" bson:"_id"`
	Params map[string]any     `json:"params" bson:"params"`
	Body   *template.Template `json:"template" bson:"template"`
}

type Attack interface {
	Template(map[string]any) (string, error)
	GenerateMail([]string, string) mail.Mail
}

func (attack *AttackType) Template(params map[string]any) (string, error) {
	buf := &bytes.Buffer{}

	if err := attack.Body.Execute(buf, params); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (attack *AttackType) validate(params map[string]any) bool {
	for k1 := range maps.Keys(attack.Params) {
		_, ok := params[k1]
		if !ok {
			return false
		}

		if reflect.TypeOf(attack.Params[k1]) != reflect.TypeOf(params[k1]) {
			return false
		}
	}

	return true
}

func (attack *AttackType) GenerateMail(params map[string]any, t []Target) (mail.Mail, error) {
	buf := &bytes.Buffer{}
	m := mail.Mail{}

	if !attack.validate(params) {
		return m, errors.New("attack has wrong params")
	}

	if err := attack.Body.Execute(buf, params); err != nil {
		return m, err
	}

	// body := buf.String()

	to := make([]string, 0)
	for value := range slices.Values(t) {
		to = append(to, value.EMail)
	}

	// m = mail.NewPlainMail(params["sender"]), params["subject"], body, to)

	return m, nil
}

//
// type AttackDownload struct {
// 	Sender       string `binding:"required" json:"sender"`
// 	Subject      string `binding:"required" json:"subject"`
// 	BodyTemplate *template.Template
// }
//
// func NewAttackDownload(body string) (AttackDownload, error) {
// 	var attackBody map[string]any
// 	var attack AttackDownload
//
// 	if err := json.Unmarshal([]byte(body), &attackBody); err != nil {
// 		return attack, err
// 	}
//
// 	t, ok := attackBody["body"].(string)
// 	if !ok {
// 		return AttackDownload{}, fmt.Errorf("failed unmarshalling attack body")
// 	}
//
// 	tmpl, err := template.New("phishingEmail").Parse(t)
// 	if err != nil {
// 		return AttackDownload{}, err
// 	}
//
// 	return AttackDownload{BodyTemplate: tmpl}, nil
// }
