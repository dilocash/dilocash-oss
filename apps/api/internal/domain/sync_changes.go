// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Syncable is a constraint for entities that support synchronization
type Syncable interface {
	GetId() uuid.UUID
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Deleted() bool
}

type SyncRepository[T Syncable] interface {
	PullChanges(context context.Context, profileId string, lastPulledAt *time.Time) (T, error)
	PushChanges(context context.Context, profileId string, lastPulledAt *time.Time, changes *T) error
}

type SyncPayload[T any] struct {
	Created []T
	Updated []T
	Deleted []uuid.UUID
}

type SyncChanges struct {
	Commands     SyncPayload[*Command]
	Intents      SyncPayload[*Intent]
	Transactions SyncPayload[*Transaction]
}
