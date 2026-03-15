// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Intent struct {
	ID             uuid.UUID
	TextMessage    string
	AudioMessage   string
	ImageMessage   string
	IntentStatus   int32
	RequiresReview bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Deleted        bool
	CommandID      uuid.UUID
}

type IntentRepository interface {
	PullChanges(context context.Context, profileId string, lastPulledAt *time.Time) (*SyncPayload[*Intent], error)
	PushChanges(context context.Context, profileId string, lastPulledAt *time.Time, intentsSync *SyncPayload[*Intent]) error
}
