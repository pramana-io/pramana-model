package model

type Resource struct {
	ID string
	Kind string
	Name string
	Provider string
	Account string
	Region string
	Tags map[string]string
}