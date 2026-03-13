// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package usecase

import (
	"context"
	"log/slog"
	"time"

	domain "github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	"golang.org/x/sync/errgroup"
)

type SyncPullUsecase struct {
	commandRepo     domain.CommandRepository
	intentRepo      domain.IntentRepository
	transactionRepo domain.TransactionRepository
}

func NewSyncPullUsecase(commandRepo domain.CommandRepository, intentRepo domain.IntentRepository, transactionRepo domain.TransactionRepository, transactor domain.Transactor) *SyncPullUsecase {
	return &SyncPullUsecase{
		commandRepo:     commandRepo,
		intentRepo:      intentRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *SyncPullUsecase) Execute(ctx context.Context, profileId string, lastPulledAt *time.Time) (*domain.SyncChanges, error) {
	slog.Debug("pulling changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	g, ctx := errgroup.WithContext(ctx)
	var syncChanges *domain.SyncChanges = &domain.SyncChanges{
		Commands:     domain.SyncPayload[*domain.Command]{},
		Intents:      domain.SyncPayload[*domain.Intent]{},
		Transactions: domain.SyncPayload[*domain.Transaction]{},
	}
	g.Go(func() error {
		commandsSync, err := u.commandRepo.PullChanges(ctx, profileId, lastPulledAt)
		if err != nil {
			return err
		}
		syncChanges.Commands = *commandsSync
		return nil
	})
	g.Go(func() error {
		intentsSync, err := u.intentRepo.PullChanges(ctx, profileId, lastPulledAt)
		if err != nil {
			return err
		}
		syncChanges.Intents = *intentsSync
		return nil
	})
	g.Go(func() error {
		transactionsSync, err := u.transactionRepo.PullChanges(ctx, profileId, lastPulledAt)
		if err != nil {
			return err
		}
		syncChanges.Transactions = *transactionsSync
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return syncChanges, nil
}
