package model

import "time"

type Snapshot struct {
	ID string
	At time.Time
	Resources map[string]Resource
	Source string
	Coverage float64
}