// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"context"
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
	PullCommandChanges(context context.Context, profileId string, lastPulledAt time.Time) (*CommandsSync, error)
	PushCommandChanges(context context.Context, profileId string, lastPulledAt time.Time, commandsSync *CommandsSync) error
}
