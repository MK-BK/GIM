package models

const (
	ScopeUser  = "USER"
	ScopeGroup = "GROUP"
)

type Scope struct {
	Kind          string
	OwnerID       string
	DestinationID string
	GroupID       string
}
