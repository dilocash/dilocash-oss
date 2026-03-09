package domain

import (
	"time"

	"github.com/google/uuid"
)

type Command struct {
	ID            uuid.UUID
	ProfileID     uuid.UUID
	CommandStatus int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Deleted       bool
}

type CommandsSync struct {
	Created []Command
	Updated []Command
	Deleted []uuid.UUID
}

type CommandRepository interface {
	PullCommandChanges(profileId string, lastPulledAt time.Time) (*CommandsSync, error)
	PushCommandChanges(profileId string, commandsSync *CommandsSync) error
}
