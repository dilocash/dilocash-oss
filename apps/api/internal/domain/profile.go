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
