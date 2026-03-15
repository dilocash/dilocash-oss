// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sync

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	v1 "github.com/dilocash/dilocash-oss/apps/api/internal/generated/transport/dilocash/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MockSyncPullUsecase is a mock of SyncPullUsecase
type MockSyncPullUsecase struct {
	mock.Mock
}

func (m *MockSyncPullUsecase) Execute(ctx context.Context, profileId string, lastPulledAt *time.Time) (*domain.SyncChanges, error) {
	args := m.Called(ctx, profileId, lastPulledAt)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.SyncChanges), args.Error(1)
}

// MockSyncPushUsecase is a mock of SyncPushUsecase
type MockSyncPushUsecase struct {
	mock.Mock
}

func (m *MockSyncPushUsecase) Execute(ctx context.Context, profileId string, lastPulledAt *time.Time, syncChanges *domain.SyncChanges) error {
	args := m.Called(ctx, profileId, lastPulledAt, syncChanges)
	return args.Error(0)
}

func TestSyncServer_PullChanges(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", "test-user")
	lastPulledAt := time.Now()
	req := &v1.PullChangesRequest{
		LastPulledAt: timestamppb.New(lastPulledAt),
	}

	t.Run("success", func(t *testing.T) {
		mockPull := new(MockSyncPullUsecase)
		mockPush := new(MockSyncPushUsecase)
		server := NewSyncServer(mockPull, mockPush)

		syncChanges := &domain.SyncChanges{
			Commands:     domain.SyncPayload[*domain.Command]{Created: []*domain.Command{{ID: uuid.New()}}},
			Intents:      domain.SyncPayload[*domain.Intent]{},
			Transactions: domain.SyncPayload[*domain.Transaction]{},
		}

		mockPull.On("Execute", mock.Anything, "test-user", mock.Anything).Return(syncChanges, nil)

		resp, err := server.PullChanges(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 1, len(resp.Changes.Commands.Created))
		mockPull.AssertExpectations(t)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockPull := new(MockSyncPullUsecase)
		mockPush := new(MockSyncPushUsecase)
		server := NewSyncServer(mockPull, mockPush)

		mockPull.On("Execute", mock.Anything, "test-user", mock.Anything).Return(nil, errors.New("pull error"))

		resp, err := server.PullChanges(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestSyncServer_PushChanges(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", "test-user")
	lastPulledAt := time.Now()
	req := &v1.PushChangesRequest{
		LastPulledAt: timestamppb.New(lastPulledAt),
		Changes: &v1.Changes{
			Commands: &v1.CommandsList{
				Created: []*v1.Command{{Id: uuid.New().String()}},
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockPull := new(MockSyncPullUsecase)
		mockPush := new(MockSyncPushUsecase)
		server := NewSyncServer(mockPull, mockPush)

		mockPush.On("Execute", mock.Anything, "test-user", mock.Anything, mock.Anything).Return(nil)

		resp, err := server.PushChanges(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.True(t, resp.Ok)
		mockPush.AssertExpectations(t)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockPull := new(MockSyncPullUsecase)
		mockPush := new(MockSyncPushUsecase)
		server := NewSyncServer(mockPull, mockPush)

		mockPush.On("Execute", mock.Anything, "test-user", mock.Anything, mock.Anything).Return(errors.New("push error"))

		resp, err := server.PushChanges(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
