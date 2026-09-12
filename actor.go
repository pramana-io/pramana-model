package model

type ActorKind string

const (
	ActorHuman ActorKind = "human"
	ActorService ActorKind = "service"
	ActorBot ActorKind = "bot"
	ActorMachine ActorKind = "machine"
	ActorUnknown ActorKind = "unknown"
)

type Actor struct {
	ID string
	Kind ActorKind
	DisplayName string
	Provider string
}