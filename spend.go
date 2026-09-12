package model

import "time"
type Money struct {
	Amount int64
	Currency string
}

type Spend struct {
	SubjectID string
	Cost Money
	PeriodStart time.Time
	PeriodEnd   time.Time
	Source string
}