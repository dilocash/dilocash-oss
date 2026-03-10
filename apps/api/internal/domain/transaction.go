// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	dv "github.com/sblackstone/shopspring-decimal-validators"
	"github.com/shopspring/decimal"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	dv.RegisterDecimalValidators(validate) // Register custom validators
}

type Transaction struct {
	ID          uuid.UUID
	Amount      decimal.Decimal `validate:"dgt=0"`
	Currency    string
	Category    string
	Description string
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required"`
	Deleted     bool
	CommandID   uuid.UUID `validate:"required"`
}

type TransactionRepository interface {
	PullChanges(context context.Context, profileId string, lastPulledAt time.Time) (*SyncPayload[*Transaction], error)
	PushChanges(context context.Context, profileId string, lastPulledAt time.Time, transactionsSync *SyncPayload[*Transaction]) error
}
