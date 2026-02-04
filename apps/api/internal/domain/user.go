// TODO add license header
package domain

import (
	"context"
	"time"
)

type User struct {
	ID                   string
	Email                string
	AcceptedTermsVersion string
	AcceptedTermsAt      *time.Time // Nullable if not yet accepted
	AllowDataAnalysis    bool       // ADR-026: Opt-in for AI training
	CreatedAt            time.Time
}

// UserRepository defines how we interact with User data.
type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	UpdateConsent(ctx context.Context, userID string, version string, allowAnalysis bool) error
}
