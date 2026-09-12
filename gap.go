package model

import "time"

type GapReason string

const (
	GapBilledNotInventoried GapReason = "billed_not_inventoried"
	GapAccountNotConnected GapReason = "account_not_connected"
	GapRegionNotQueried GapReason = "region_not_queried"
	GapOwnerUnknown GapReason = "owner_unknown"
	GapAutomationUnbounded GapReason = "automation_unbounded"
)

type Gap struct {
	SubjectID string
	Reason GapReason
	DetectedAt time.Time
	Source string
	Detail string
}