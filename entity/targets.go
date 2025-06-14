package entity

type TargetState int

const (
	StateActive TargetState = iota
	StateCompleted
	StateError
	StateRetrying
)

var stateName = map[TargetState]string{
	StateActive:    "active",
	StateCompleted: "completed",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ts TargetState) String() string {
	return stateName[ts]
}

type Target struct {
	ID             string      `binding:"required" rethinkdb:"id" json:"id"`
	OrganizationID string      `binding:"required" rethinkdb:"organization_id" json:"organization_id"`
	EMail          string      `binding:"required" rethinkdb:"e_mail" json:"e_mail"`
	Firstname      string      `binding:"required" rethinkdb:"firstname" json:"firstname"`
	Surname        string      `binding:"required" rethinkdb:"surname" json:"surname"`
	State          TargetState `binding:"required" rethinkdb:"state" json:"state"`
}
