// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Amount      decimal.Decimal
	Currency    string
	Category    string
	Description string
	RawInput    string
	CreatedAt   time.Time
}

type User struct {
	ID                   uuid.UUID
	Email                string
	AcceptedTermsVersion string
	AcceptedTermsAt      time.Time
	AllowDataAnalysis    bool
	CreatedAt            time.Time
}
