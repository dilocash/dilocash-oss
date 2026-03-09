package usecase

import (
	"context"
	"log/slog"
	"time"

	domain "github.com/dilocash/dilocash-oss/apps/api/internal/domain"
)

type SyncPullUsecase struct {
	commandRepo     domain.CommandRepository
	intentRepo      domain.IntentRepository
	transactionRepo domain.TransactionRepository
}

func NewSyncPullUsecase(commandRepo domain.CommandRepository, intentRepo domain.IntentRepository, transactionRepo domain.TransactionRepository) *SyncPullUsecase {
	return &SyncPullUsecase{
		commandRepo:     commandRepo,
		intentRepo:      intentRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *SyncPullUsecase) Execute(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.SyncChanges, error) {
	slog.Info("pulling changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	commandsSync, err := u.commandRepo.PullCommandChanges(profileId, lastPulledAt)
	if err != nil {
		return nil, err
	}
	intentsSync, err := u.intentRepo.PullIntentChanges(profileId, lastPulledAt)
	if err != nil {
		return nil, err
	}
	transactionsSync, err := u.transactionRepo.PullTransactionChanges(profileId, lastPulledAt)
	if err != nil {
		return nil, err
	}
	return &domain.SyncChanges{
		Commands:     *commandsSync,
		Intents:      *intentsSync,
		Transactions: *transactionsSync,
	}, nil
}
