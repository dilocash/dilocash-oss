// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID                   uuid.UUID
	DisplayName          string
	Email                string
	AcceptedTermsVersion string
	AcceptedTermsAt      time.Time
	AllowDataAnalysis    bool
	CreatedAt            time.Time
}
