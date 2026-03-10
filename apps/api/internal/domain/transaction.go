// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          uuid.UUID
	Amount      decimal.Decimal
	Currency    string
	Category    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
	CommandID   uuid.UUID
}

type TransactionRepository interface {
	PullChanges(context context.Context, profileId string, lastPulledAt time.Time) (*SyncPayload[*Transaction], error)
	PushChanges(context context.Context, profileId string, lastPulledAt time.Time, transactionsSync *SyncPayload[*Transaction]) error
}
