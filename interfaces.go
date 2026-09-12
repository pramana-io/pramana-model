package model

import (
	"context"
	"time"
)

type Inventory interface {
	List(ctx context.Context) ([]Resource, error)
}

type Oracle interface {
	Costs(ctx context.Context, start, end time.Time) ([]Spend, error)
}

type Reconciler interface {
	Reconcile(ctx context.Context, snap Snapshot, costs []Spend) ([]Gap, error)
}

type Store interface {
	PutSnapshot(ctx context.Context, snap Snapshot) error
	PutAction(ctx context.Context, act Action) error
	SnapshotAt(ctx context.Context, at time.Time) (Snapshot, error)
}