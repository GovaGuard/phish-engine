package entity

import "time"

type Campaign struct {
	ID             string    `binding:"required" json:"id" rethinkdb:"id"`
	CreatorID      string    `binding:"required" json:"creator_id" rethinkdb:"creator_id"`
	OrganizationID string    `binding:"required" json:"organization_id" rethinkdb:"organization_id"`
	Title          string    `binding:"required" json:"title" rethinkdb:"title"`
	Status         string    `binding:"required" json:"status" rethinkdb:"status"`
	StartDate      time.Time `binding:"required" json:"start_date" rethinkdb:"start_date"`
	EndDate        time.Time `binding:"required" json:"end_date" rethinkdb:"end_date"`
	SuccessRate    int16     `binding:"required" json:"success_rate" rethinkdb:"success_rate"`
	Targets        []Target  `binding:"required" json:"targets" rethinkdb:"targets"`
}
