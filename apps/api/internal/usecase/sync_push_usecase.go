// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package usecase

import (
	"context"
	"log/slog"
	"time"

	domain "github.com/dilocash/dilocash-oss/apps/api/internal/domain"
)

type SyncPushUsecase struct {
	commandRepo     domain.CommandRepository
	intentRepo      domain.IntentRepository
	transactionRepo domain.TransactionRepository
	transactor      domain.Transactor
}

func NewSyncPushUsecase(commandRepo domain.CommandRepository, intentRepo domain.IntentRepository, transactionRepo domain.TransactionRepository, transactor domain.Transactor) *SyncPushUsecase {
	return &SyncPushUsecase{
		commandRepo:     commandRepo,
		intentRepo:      intentRepo,
		transactionRepo: transactionRepo,
		transactor:      transactor,
	}
}

func (u *SyncPushUsecase) Execute(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.SyncChanges, error) {
	slog.Info("pushing changes", "profileId", profileId, "lastPulledAt", lastPulledAt)

	var syncChanges *domain.SyncChanges

	// El transactor inyecta la TX en el ctx y se lo pasa al repo
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		commandsSync, err := u.commandRepo.PullCommandChanges(ctx, profileId, lastPulledAt)
		if err != nil {
			return err
		}
		intentsSync, err := u.intentRepo.PullIntentChanges(ctx, profileId, lastPulledAt)
		if err != nil {
			return err
		}
		transactionsSync, err := u.transactionRepo.PullTransactionChanges(ctx, profileId, lastPulledAt)
		if err != nil {
			return err
		}
		syncChanges = &domain.SyncChanges{
			Commands:     *commandsSync,
			Intents:      *intentsSync,
			Transactions: *transactionsSync,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return syncChanges, nil
}
