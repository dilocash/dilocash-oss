// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSyncPullUsecase_Execute(t *testing.T) {
	ctx := context.Background()
	profileId := "test-profile-id"
	lastPulledAt := time.Now()

	t.Run("success", func(t *testing.T) {
		mockCommandRepo := new(MockCommandRepository)
		mockIntentRepo := new(MockIntentRepository)
		mockTransactionRepo := new(MockTransactionRepository)

		commands := &domain.SyncPayload[*domain.Command]{
			Created: []*domain.Command{{ID: uuid.New()}},
			Updated: []*domain.Command{},
			Deleted: []uuid.UUID{uuid.New()},
		}
		intents := &domain.SyncPayload[*domain.Intent]{
			Created: []*domain.Intent{{ID: uuid.New()}},
			Updated: []*domain.Intent{},
			Deleted: []uuid.UUID{},
		}
		transactions := &domain.SyncPayload[*domain.Transaction]{
			Created: []*domain.Transaction{{ID: uuid.New()}},
			Updated: []*domain.Transaction{},
			Deleted: []uuid.UUID{},
		}

		mockCommandRepo.On("PullChanges", mock.Anything, profileId, lastPulledAt).Return(commands, nil)
		mockIntentRepo.On("PullChanges", mock.Anything, profileId, lastPulledAt).Return(intents, nil)
		mockTransactionRepo.On("PullChanges", mock.Anything, profileId, lastPulledAt).Return(transactions, nil)

		u := NewSyncPullUsecase(mockCommandRepo, mockIntentRepo, mockTransactionRepo, nil)
		result, err := u.Execute(ctx, profileId, lastPulledAt)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *commands, result.Commands)
		assert.Equal(t, *intents, result.Intents)
		assert.Equal(t, *transactions, result.Transactions)

		mockCommandRepo.AssertExpectations(t)
		mockIntentRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		mockCommandRepo := new(MockCommandRepository)
		mockIntentRepo := new(MockIntentRepository)
		mockTransactionRepo := new(MockTransactionRepository)

		mockCommandRepo.On("PullChanges", mock.Anything, profileId, lastPulledAt).Return(nil, errors.New("repo error"))
		mockIntentRepo.On("PullChanges", mock.Anything, profileId, lastPulledAt).Return(&domain.SyncPayload[*domain.Intent]{}, nil)
		mockTransactionRepo.On("PullChanges", mock.Anything, profileId, lastPulledAt).Return(&domain.SyncPayload[*domain.Transaction]{}, nil)

		u := NewSyncPullUsecase(mockCommandRepo, mockIntentRepo, mockTransactionRepo, nil)
		result, err := u.Execute(ctx, profileId, lastPulledAt)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "repo error", err.Error())
	})
}
