package domain

import (
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

type IntentsSync struct {
	Created []Intent
	Updated []Intent
	Deleted []uuid.UUID
}

type IntentRepository interface {
	PullIntentChanges(profileId string, lastPulledAt time.Time) (*IntentsSync, error)
	PushIntentChanges(profileId string, intentsSync *IntentsSync) error
}
