// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDomainStructs(t *testing.T) {
	t.Run("Command", func(t *testing.T) {
		id := uuid.New()
		profileId := uuid.New()
		now := time.Now()
		cmd := &Command{
			ID:            id,
			ProfileID:     profileId,
			CommandStatus: 1,
			CreatedAt:     now,
			UpdatedAt:     now,
			Deleted:       false,
		}
		assert.Equal(t, id, cmd.ID)
		assert.Equal(t, profileId, cmd.ProfileID)
		assert.Equal(t, now, cmd.CreatedAt)
	})

	t.Run("Intent", func(t *testing.T) {
		id := uuid.New()
		intent := &Intent{
			ID:           id,
			TextMessage:  "test",
			IntentStatus: 1,
		}
		assert.Equal(t, id, intent.ID)
		assert.Equal(t, "test", intent.TextMessage)
	})

	TestTransactionDomainStructs(t)
}
