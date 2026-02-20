// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Command struct {
	ID            uuid.UUID
	ProfileID     uuid.UUID
	CommandStatus int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Deleted       bool
}

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

type Profile struct {
	ID                   uuid.UUID
	DisplayName          string
	Email                string
	AcceptedTermsVersion string
	AcceptedTermsAt      time.Time
	AllowDataAnalysis    bool
	CreatedAt            time.Time
}
