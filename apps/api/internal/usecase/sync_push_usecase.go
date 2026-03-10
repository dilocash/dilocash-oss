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

func (u *SyncPushUsecase) Execute(ctx context.Context, profileId string, lastPulledAt time.Time, syncChanges *domain.SyncChanges) error {
	slog.Info("pushing changes", "profileId", profileId, "lastPulledAt", lastPulledAt)

	// El transactor inyecta la TX en el ctx y se lo pasa al repo
	err := u.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		err := u.commandRepo.PushCommandChanges(ctx, profileId, lastPulledAt, &syncChanges.Commands)
		if err != nil {
			return err
		}
		err = u.intentRepo.PushIntentChanges(ctx, profileId, lastPulledAt, &syncChanges.Intents)
		if err != nil {
			return err
		}
		err = u.transactionRepo.PushTransactionChanges(ctx, profileId, lastPulledAt, &syncChanges.Transactions)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
