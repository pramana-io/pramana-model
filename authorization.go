package model

type Permission string

type Authorization struct {
	ActorID string
	Granted []Permission
	Used []Permission
}