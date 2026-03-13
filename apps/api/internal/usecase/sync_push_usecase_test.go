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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSyncPushUsecase_Execute(t *testing.T) {
	ctx := context.Background()
	profileId := "test-profile-id"
	lastPulledAt := time.Now()

	t.Run("success", func(t *testing.T) {
		mockCommandRepo := new(MockCommandRepository)
		mockIntentRepo := new(MockIntentRepository)
		mockTransactionRepo := new(MockTransactionRepository)
		mockTransactor := new(MockTransactor)

		syncChanges := &domain.SyncChanges{
			Commands:     domain.SyncPayload[*domain.Command]{},
			Intents:      domain.SyncPayload[*domain.Intent]{},
			Transactions: domain.SyncPayload[*domain.Transaction]{},
		}

		mockTransactor.On("WithinTransaction", ctx, mock.AnythingOfType("func(context.Context) error")).Return(nil)
		mockCommandRepo.On("PushChanges", mock.Anything, profileId, lastPulledAt, &syncChanges.Commands).Return(nil)
		mockIntentRepo.On("PushChanges", mock.Anything, profileId, lastPulledAt, &syncChanges.Intents).Return(nil)
		mockTransactionRepo.On("PushChanges", mock.Anything, profileId, lastPulledAt, &syncChanges.Transactions).Return(nil)

		u := NewSyncPushUsecase(mockCommandRepo, mockIntentRepo, mockTransactionRepo, mockTransactor)
		err := u.Execute(ctx, profileId, &lastPulledAt, syncChanges)

		assert.NoError(t, err)

		mockTransactor.AssertExpectations(t)
		mockCommandRepo.AssertExpectations(t)
		mockIntentRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("transactor error", func(t *testing.T) {
		mockCommandRepo := new(MockCommandRepository)
		mockIntentRepo := new(MockIntentRepository)
		mockTransactionRepo := new(MockTransactionRepository)
		mockTransactor := new(MockTransactor)

		mockTransactor.On("WithinTransaction", ctx, mock.AnythingOfType("func(context.Context) error")).Return(errors.New("tx error"))

		lastPulledAt := time.Now()
		u := NewSyncPushUsecase(mockCommandRepo, mockIntentRepo, mockTransactionRepo, mockTransactor)
		err := u.Execute(ctx, "id", &lastPulledAt, &domain.SyncChanges{})

		assert.Error(t, err)
		assert.Equal(t, "tx error", err.Error())
	})

	t.Run("repository error inside transaction", func(t *testing.T) {
		mockCommandRepo := new(MockCommandRepository)
		mockIntentRepo := new(MockIntentRepository)
		mockTransactionRepo := new(MockTransactionRepository)
		mockTransactor := new(MockTransactor)

		syncChanges := &domain.SyncChanges{}

		mockTransactor.On("WithinTransaction", ctx, mock.AnythingOfType("func(context.Context) error")).Return(errors.New("repo error"))
		mockCommandRepo.On("PushChanges", mock.Anything, profileId, lastPulledAt, &syncChanges.Commands).Return(errors.New("repo error"))

		u := NewSyncPushUsecase(mockCommandRepo, mockIntentRepo, mockTransactionRepo, mockTransactor)
		err := u.Execute(ctx, profileId, &lastPulledAt, syncChanges)

		assert.Error(t, err)
		assert.Equal(t, "repo error", err.Error())
	})
}
