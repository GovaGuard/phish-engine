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
	ID             string      `bson:"id" json:"id"`
	OrganizationID string      `bson:"organization_id" json:"organization_id"`
	EMail          string      `binding:"required" bson:"e_mail" json:"e_mail"`
	Firstname      string      `binding:"required" bson:"firstname" json:"firstname"`
	Surname        string      `binding:"required" bson:"surname" json:"surname"`
	State          TargetState `binding:"required" bson:"state" json:"state"`
}
