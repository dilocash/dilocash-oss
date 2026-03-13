// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package usecase

import (
	"context"
	"time"

	"github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockCommandRepository is a mock of CommandRepository
type MockCommandRepository struct {
	mock.Mock
}

func (m *MockCommandRepository) PullChanges(ctx context.Context, profileId string, lastPulledAt *time.Time) (*domain.SyncPayload[*domain.Command], error) {
	args := m.Called(ctx, profileId, lastPulledAt)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.SyncPayload[*domain.Command]), args.Error(1)
}

func (m *MockCommandRepository) PushChanges(ctx context.Context, profileId string, lastPulledAt *time.Time, commandsSync *domain.SyncPayload[*domain.Command]) error {
	args := m.Called(ctx, profileId, lastPulledAt, commandsSync)
	return args.Error(0)
}

// MockIntentRepository is a mock of IntentRepository
type MockIntentRepository struct {
	mock.Mock
}

func (m *MockIntentRepository) PullChanges(ctx context.Context, profileId string, lastPulledAt *time.Time) (*domain.SyncPayload[*domain.Intent], error) {
	args := m.Called(ctx, profileId, lastPulledAt)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.SyncPayload[*domain.Intent]), args.Error(1)
}

func (m *MockIntentRepository) PushChanges(ctx context.Context, profileId string, lastPulledAt *time.Time, intentsSync *domain.SyncPayload[*domain.Intent]) error {
	args := m.Called(ctx, profileId, lastPulledAt, intentsSync)
	return args.Error(0)
}

// MockTransactionRepository is a mock of TransactionRepository
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) PullChanges(ctx context.Context, profileId string, lastPulledAt *time.Time) (*domain.SyncPayload[*domain.Transaction], error) {
	args := m.Called(ctx, profileId, lastPulledAt)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.SyncPayload[*domain.Transaction]), args.Error(1)
}

func (m *MockTransactionRepository) PushChanges(ctx context.Context, profileId string, lastPulledAt *time.Time, transactionsSync *domain.SyncPayload[*domain.Transaction]) error {
	args := m.Called(ctx, profileId, lastPulledAt, transactionsSync)
	return args.Error(0)
}

// MockTransactor is a mock of Transactor
type MockTransactor struct {
	mock.Mock
}

func (m *MockTransactor) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	args := m.Called(ctx, fn)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return fn(ctx)
}
