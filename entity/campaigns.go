package entity

type Campaign struct {
	ID             string   `json:"id" rethinkdb:"id"`
	CreatorID      string   `json:"creator_id" rethinkdb:"creator_id"`
	OrganizationID string   `json:"organization_id" rethinkdb:"organization_id"`
	Title          string   `json:"title" rethinkdb:"title"`
	State          string   `json:"state" rethinkdb:"state"`
	StartDate      string   `json:"start_date" rethinkdb:"start_date"`
	EndDate        string   `json:"end_date" rethinkdb:"end_date"`
	Targets        []Target `json:"targets" rethinkdb:"targets"`
}

type Target struct {
	ID             string
	OrganizationID string
	EMail          string
	FirstName      string
	SurName        string
	State          string
}
