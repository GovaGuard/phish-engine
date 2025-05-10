package entity

type Target struct {
	ID             string `binding:"required" rethinkdb:"id" json:"id"`
	OrganizationID string `binding:"required" rethinkdb:"organization_id" json:"organization_id"`
	EMail          string `binding:"required" rethinkdb:"e_mail" json:"e_mail"`
	Firstname      string `binding:"required" rethinkdb:"firstname" json:"firstname"`
	Surname        string `binding:"required" rethinkdb:"surname" json:"surname"`
	State          string `binding:"required" rethinkdb:"state" json:"state"`
}
