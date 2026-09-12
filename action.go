package model

import "time"

type ActionKind string

const (
	ActionCreate ActionKind = "create"
	ActionRead ActionKind = "read"
	ActionUpdate ActionKind = "update"
	ActionDelete ActionKind = "delete"
	ActionUnknown ActionKind = "unknown"
)

type Action struct {
	ID string
	Kind ActionKind
	ActorID string
	ResourceID string
	At time.Time
	Authorized bool
	Source string
}