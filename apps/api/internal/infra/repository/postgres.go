// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"connectrpc.com/connect"
	"github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	db "github.com/dilocash/dilocash-oss/apps/api/internal/generated/db/postgres"
	mappers "github.com/dilocash/dilocash-oss/apps/api/internal/generated/mappers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type PostgresRepo struct {
	q    *db.Queries
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		q:    db.New(pool),
		pool: pool,
	}
}

func (r *PostgresRepo) PullCommandChanges(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.CommandsSync, error) {
	converter := &mappers.ConverterImpl{}
	created := []domain.Command{}

	executor := r.getDB(ctx)
	q := db.New(executor)

	slog.Info("pulling command changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	// execute query to get created commands since lastPulledAt
	createdCommandsResult, err := q.ListCommandsByProfileIdAndCreatedAfter(ctx, db.ListCommandsByProfileIdAndCreatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		CreatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull created commands", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull created commands"))
	}

	for _, dbCommandResult := range createdCommandsResult {
		dbCommand := dbCommandResult
		domainCommand := converter.CommandFromDBToDomain(dbCommand)
		created = append(created, domainCommand)
	}

	updated := []domain.Command{}
	// execute query to get updated commands since lastPulledAt
	updatedCommandsResult, err := q.ListCommandsByProfileIdAndUpdatedAfter(ctx, db.ListCommandsByProfileIdAndUpdatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		UpdatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull updated commands", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull updated commands"))
	}

	for _, dbCommandResult := range updatedCommandsResult {
		dbCommand := dbCommandResult
		domainCommand := converter.CommandFromDBToDomain(dbCommand)
		updated = append(updated, domainCommand)
	}

	deleted := []uuid.UUID{}

	// execute query to get deleted commands since lastPulledAt
	deletedCommandsResult, err := q.ListDeletedCommandsByProfileIdAndUpdatedAfter(ctx, db.ListDeletedCommandsByProfileIdAndUpdatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		UpdatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull deleted commands", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull deleted commands"))
	}

	for _, dbCommandResult := range deletedCommandsResult {
		deleted = append(deleted, dbCommandResult)
	}

	return &domain.CommandsSync{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}, nil
}

func (r *PostgresRepo) PullIntentChanges(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.IntentsSync, error) {
	converter := &mappers.ConverterImpl{}
	created := []domain.Intent{}

	executor := r.getDB(ctx)
	q := db.New(executor)
	slog.Info("pulling intent changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	// execute query to get intents since lastPulledAt
	createdIntentsResult, err := q.ListIntentsByProfileIdAndCreatedAfter(ctx, db.ListIntentsByProfileIdAndCreatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		CreatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull created intents", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull created intents"))
	}

	for _, dbIntentResult := range createdIntentsResult {
		dbIntent := dbIntentResult
		domainIntent := converter.IntentFromDBToDomain(dbIntent)
		created = append(created, domainIntent)
	}

	updated := []domain.Intent{}
	// execute query to get updated intents since lastPulledAt
	updatedIntentsResult, err := q.ListIntentsByProfileIdAndUpdatedAfter(ctx, db.ListIntentsByProfileIdAndUpdatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		UpdatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull updated intents", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull updated intents"))
	}

	for _, dbIntentResult := range updatedIntentsResult {
		dbIntent := dbIntentResult
		domainIntent := converter.IntentFromDBToDomain(dbIntent)
		updated = append(updated, domainIntent)
	}

	deleted := []uuid.UUID{}
	// execute query to get deleted intents since lastPulledAt
	deletedIntentsResult, err := q.ListDeletedIntentsByProfileIdAndUpdatedAfter(ctx, db.ListDeletedIntentsByProfileIdAndUpdatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		UpdatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull deleted intents", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull deleted intents"))
	}

	for _, dbIntentResult := range deletedIntentsResult {
		deleted = append(deleted, dbIntentResult)
	}
	return &domain.IntentsSync{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}, nil
}

func (r *PostgresRepo) PullTransactionChanges(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.TransactionsSync, error) {
	converter := &mappers.ConverterImpl{}
	created := []domain.Transaction{}

	executor := r.getDB(ctx)
	q := db.New(executor)
	slog.Info("pulling transaction changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	// execute query to get transactions since lastPulledAt
	createdTransactionsResult, err := q.ListTransactionsByProfileIdAndCreatedAfter(ctx, db.ListTransactionsByProfileIdAndCreatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		CreatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull created transactions", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull created transactions"))
	}

	for _, dbTransactionResult := range createdTransactionsResult {
		dbTransaction := dbTransactionResult
		domainTransaction := converter.TransactionFromDBToDomain(dbTransaction)
		created = append(created, domainTransaction)
	}

	updated := []domain.Transaction{}
	// execute query to get updated transactions since lastPulledAt
	updatedTransactionsResult, err := q.ListTransactionsByProfileIdAndUpdatedAfter(ctx, db.ListTransactionsByProfileIdAndUpdatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		UpdatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull updated transactions", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull updated transactions"))
	}

	for _, dbTransactionResult := range updatedTransactionsResult {
		dbTransaction := dbTransactionResult
		domainTransaction := converter.TransactionFromDBToDomain(dbTransaction)
		updated = append(updated, domainTransaction)
	}

	deleted := []uuid.UUID{}
	// execute query to get deleted transactions since lastPulledAt
	deletedTransactionsResult, err := q.ListDeletedTransactionsByProfileIdAndUpdatedAfter(ctx, db.ListDeletedTransactionsByProfileIdAndUpdatedAfterParams{
		ProfileID: uuid.MustParse(profileId),
		UpdatedAt: lastPulledAt,
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		slog.Error("failed to pull deleted transactions", "error", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to pull deleted transactions"))
	}

	for _, dbTransactionResult := range deletedTransactionsResult {
		deleted = append(deleted, dbTransactionResult)
	}
	return &domain.TransactionsSync{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}, nil
}

func (r *PostgresRepo) PushCommandChanges(ctx context.Context, profileId string, lastPulledAt time.Time, commandsSync *domain.CommandsSync) error {
	executor := r.getDB(ctx)
	q := db.New(executor)
	for _, command := range commandsSync.Created {
		params := db.CreateCommandParams{
			ID:            command.ID,
			ProfileID:     uuid.MustParse(profileId),
			CommandStatus: command.CommandStatus,
			CreatedAt:     command.CreatedAt,
		}
		_, err := q.CreateCommand(ctx, params)
		if err != nil {
			slog.Error("failed to store command", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store command"))
		}
	}

	for _, command := range commandsSync.Updated {
		params := db.UpdateCommandParams{
			ID:            command.ID,
			CommandStatus: command.CommandStatus,
		}
		_, err := q.UpdateCommand(ctx, params)
		if err != nil {
			slog.Error("failed to store command", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store command"))
		}
	}

	for _, command := range commandsSync.Deleted {
		err := q.DeleteCommand(ctx, command)
		if err != nil {
			slog.Error("failed to delete command", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to delete command"))
		}
	}
	return nil
}

func (r *PostgresRepo) PushIntentChanges(ctx context.Context, profileId string, lastPulledAt time.Time, intentsSync *domain.IntentsSync) error {
	executor := r.getDB(ctx)
	q := db.New(executor)
	for _, intent := range intentsSync.Created {
		params := db.CreateIntentParams{
			ID:             intent.ID,
			CommandID:      intent.CommandID,
			TextMessage:    &intent.TextMessage,
			AudioMessage:   &intent.AudioMessage,
			ImageMessage:   &intent.ImageMessage,
			IntentStatus:   intent.IntentStatus,
			RequiresReview: &intent.RequiresReview,
			CreatedAt:      intent.CreatedAt,
		}
		_, err := q.CreateIntent(ctx, params)
		if err != nil {
			slog.Error("failed to store intent", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store intent"))
		}
	}

	for _, intent := range intentsSync.Updated {
		params := db.UpdateIntentParams{
			ID:           intent.ID,
			IntentStatus: intent.IntentStatus,
		}
		_, err := q.UpdateIntent(ctx, params)
		if err != nil {
			slog.Error("failed to store intent", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store intent"))
		}
	}

	for _, intent := range intentsSync.Deleted {
		err := q.DeleteIntent(ctx, intent)
		if err != nil {
			slog.Error("failed to delete intent", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to delete intent"))
		}
	}
	return nil
}

func (r *PostgresRepo) PushTransactionChanges(ctx context.Context, profileId string, lastPulledAt time.Time, transactionsSync *domain.TransactionsSync) error {
	executor := r.getDB(ctx)
	q := db.New(executor)

	for _, intent := range transactionsSync.Created {
		params := db.CreateTransactionParams{
			ID:          intent.ID,
			CommandID:   intent.CommandID,
			Amount:      decimal.RequireFromString(intent.Amount.String()),
			Currency:    intent.Currency,
			Category:    &intent.Category,
			Description: &intent.Description,
			CreatedAt:   intent.CreatedAt,
		}
		_, err := q.CreateTransaction(ctx, params)
		if err != nil {
			slog.Error("failed to store transaction", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store transaction"))
		}
	}

	for _, transaction := range transactionsSync.Updated {
		params := db.UpdateTransactionParams{
			ID:          transaction.ID,
			Amount:      decimal.RequireFromString(transaction.Amount.String()),
			Currency:    transaction.Currency,
			Category:    &transaction.Category,
			Description: &transaction.Description,
		}
		_, err := q.UpdateTransaction(ctx, params)
		if err != nil {
			slog.Error("failed to store transaction", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store transaction"))
		}
	}

	for _, intent := range transactionsSync.Deleted {
		err := q.DeleteTransaction(ctx, intent)
		if err != nil {
			slog.Error("failed to delete transaction", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to delete transaction"))
		}
	}
	return nil
}

func (r *PostgresRepo) getDB(ctx context.Context) db.DBTX {
	if tx, ok := extractTx(ctx); ok {
		return tx
	}
	return r.pool
}
