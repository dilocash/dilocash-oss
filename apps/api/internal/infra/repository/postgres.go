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
)

// could be used in the future for repos which access database but are not syncable
type PostgresRepo struct {
	q    *db.Queries
	pool *pgxpool.Pool
}

// BaseSyncRepo handles the common logic for pgx/sqlc
type BaseSyncRepo[DBEntity any, DomainEntity any] struct {
	q         *db.Queries
	pool      *pgxpool.Pool
	toDomain  func(DBEntity) DomainEntity
	toDB      func(DomainEntity) DBEntity
	converter *mappers.ConverterImpl
}

type CommandRepository struct {
	*BaseSyncRepo[db.Command, *domain.Command]
}

type IntentRepository struct {
	*BaseSyncRepo[db.Intent, *domain.Intent]
}

type TransactionRepository struct {
	*BaseSyncRepo[db.Transaction, *domain.Transaction]
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		q:    db.New(pool),
		pool: pool,
	}
}

func NewCommandRepository(pool *pgxpool.Pool, conv *mappers.ConverterImpl) *CommandRepository {
	return &CommandRepository{
		BaseSyncRepo: &BaseSyncRepo[db.Command, *domain.Command]{
			pool:      pool,
			q:         db.New(pool),
			toDomain:  conv.CommandFromDBToDomain,
			toDB:      conv.ToDBCommand,
			converter: conv,
		},
	}
}

func NewIntentRepository(pool *pgxpool.Pool, conv *mappers.ConverterImpl) *IntentRepository {
	return &IntentRepository{
		BaseSyncRepo: &BaseSyncRepo[db.Intent, *domain.Intent]{
			pool:      pool,
			q:         db.New(pool),
			toDomain:  conv.IntentFromDBToDomain,
			toDB:      conv.ToDBIntent,
			converter: conv,
		},
	}
}

func NewTransactionRepository(pool *pgxpool.Pool, conv *mappers.ConverterImpl) *TransactionRepository {
	return &TransactionRepository{
		BaseSyncRepo: &BaseSyncRepo[db.Transaction, *domain.Transaction]{
			pool:      pool,
			q:         db.New(pool),
			toDomain:  conv.TransactionFromDBToDomain,
			toDB:      conv.ToDBTransaction,
			converter: conv,
		},
	}
}

/*
Generic PullChanges which queries created, updated and deleted records for a Dom entity
*/
func (r *BaseSyncRepo[DB, Dom]) PullChanges(ctx context.Context, profileId string, lastPulledAt time.Time, fetch func(ctx context.Context, updatedAfter time.Time) ([]DB, []DB, []uuid.UUID, error)) (*domain.SyncPayload[Dom], error) {
	// execute queries to get created, updated and deleted entities since lastPulledAt
	createdRows, updatedRows, deletedRows, err := fetch(ctx, lastPulledAt)
	if err != nil {
		return nil, err
	}
	// convert db entities to domain entities
	created := make([]Dom, len(createdRows))
	for i, row := range createdRows {
		created[i] = r.toDomain(row)
	}

	// convert db entities to domain entities
	updated := make([]Dom, len(updatedRows))
	for i, row := range updatedRows {
		updated[i] = r.toDomain(row)
	}

	// deleted entities are already uuid, no convertion needed
	//return created, updated and deletedRows
	return &domain.SyncPayload[Dom]{
		Created: created,
		Updated: updated,
		Deleted: deletedRows,
	}, nil
}

func (r *CommandRepository) PullChanges(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.SyncPayload[*domain.Command], error) {
	// we call generic method with specific sqlc generated queries
	return r.BaseSyncRepo.PullChanges(ctx, profileId, lastPulledAt, func(ctx context.Context, updatedAfter time.Time) ([]db.Command, []db.Command, []uuid.UUID, error) {
		executor := r.getDB(ctx)
		q := db.New(executor)
		created := []db.Command{}
		updated := []db.Command{}
		deleted := []uuid.UUID{}
		slog.Debug("querying sync commands", "profileId", profileId, "lastPulledAt", lastPulledAt)
		rows, err := q.GetCommandsSync(ctx, db.GetCommandsSyncParams{
			ProfileID: uuid.MustParse(profileId),
			CreatedAt: lastPulledAt,
			Limit:     100,
			Offset:    0,
		})
		if err != nil {
			slog.Error("failed to query sync commands", "error", err)
			return nil, nil, nil, connect.NewError(connect.CodeInternal, errors.New("failed to query sync commands"))
		}
		for _, row := range rows {
			switch row.SyncType {
			case "deleted":
				deleted = append(deleted, row.ID)
			case "created":
				created = append(created, r.converter.CommandRowFromDBToDB(row))
			case "updated":
				updated = append(updated, r.converter.CommandRowFromDBToDB(row))
			}
		}

		return created, updated, deleted, nil
	})
}

func (r *CommandRepository) PushChanges(ctx context.Context, profileId string, lastPulledAt time.Time, commandsSync *domain.SyncPayload[*domain.Command]) error {
	executor := r.getDB(ctx)
	q := db.New(executor)
	for _, command := range commandsSync.Created {
		params := r.converter.ToDBCreateCommandParams(command)
		_, err := q.CreateCommand(ctx, params)
		if err != nil {
			slog.Error("failed to store command", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store command"))
		}
	}

	for _, command := range commandsSync.Updated {
		params := r.converter.ToDBUpdateCommandParams(command)
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

func (r *IntentRepository) PullChanges(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.SyncPayload[*domain.Intent], error) {
	// we call generic method with specific sqlc generated queries
	return r.BaseSyncRepo.PullChanges(ctx, profileId, lastPulledAt, func(ctx context.Context, updatedAfter time.Time) ([]db.Intent, []db.Intent, []uuid.UUID, error) {
		executor := r.getDB(ctx)
		q := db.New(executor)
		created := []db.Intent{}
		updated := []db.Intent{}
		deleted := []uuid.UUID{}
		slog.Debug("querying sync intents", "profileId", profileId, "lastPulledAt", lastPulledAt)
		rows, err := q.GetIntentsSync(ctx, db.GetIntentsSyncParams{
			ProfileID: uuid.MustParse(profileId),
			CreatedAt: lastPulledAt,
			Limit:     100,
			Offset:    0,
		})
		if err != nil {
			slog.Error("failed to query sync intents", "error", err)
			return nil, nil, nil, connect.NewError(connect.CodeInternal, errors.New("failed to query sync intents"))
		}
		for _, row := range rows {
			switch row.SyncType {
			case "deleted":
				deleted = append(deleted, row.ID)
			case "created":
				created = append(created, r.converter.IntentRowFromDBToDB(row))
			case "updated":
				updated = append(updated, r.converter.IntentRowFromDBToDB(row))
			}
		}
		return created, updated, deleted, nil
	})
}

func (r *TransactionRepository) PullChanges(ctx context.Context, profileId string, lastPulledAt time.Time) (*domain.SyncPayload[*domain.Transaction], error) {
	// we call generic method with specific sqlc generated queries
	return r.BaseSyncRepo.PullChanges(ctx, profileId, lastPulledAt, func(ctx context.Context, updatedAfter time.Time) ([]db.Transaction, []db.Transaction, []uuid.UUID, error) {
		executor := r.getDB(ctx)
		q := db.New(executor)
		created := []db.Transaction{}
		updated := []db.Transaction{}
		deleted := []uuid.UUID{}
		slog.Debug("querying sync transactions", "profileId", profileId, "lastPulledAt", lastPulledAt)
		rows, err := q.GetTransactionsSync(ctx, db.GetTransactionsSyncParams{
			ProfileID: uuid.MustParse(profileId),
			CreatedAt: lastPulledAt,
			Limit:     100,
			Offset:    0,
		})
		if err != nil {
			slog.Error("failed to query sync transactions", "error", err)
			return nil, nil, nil, connect.NewError(connect.CodeInternal, errors.New("failed to query sync transactions"))
		}
		for _, row := range rows {
			switch row.SyncType {
			case "deleted":
				deleted = append(deleted, row.ID)
			case "created":
				created = append(created, r.converter.TransactionRowFromDBToDB(row))
			case "updated":
				updated = append(updated, r.converter.TransactionRowFromDBToDB(row))
			}
		}

		return created, updated, deleted, nil
	})
}

func (r *IntentRepository) PushChanges(ctx context.Context, profileId string, lastPulledAt time.Time, intentsSync *domain.SyncPayload[*domain.Intent]) error {
	executor := r.getDB(ctx)
	q := db.New(executor)
	for _, intent := range intentsSync.Created {
		params := r.converter.ToDBCreateIntentParams(intent)
		_, err := q.CreateIntent(ctx, params)
		if err != nil {
			slog.Error("failed to store intent", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store intent"))
		}
	}

	for _, intent := range intentsSync.Updated {
		params := r.converter.ToDBUpdateIntentParams(intent)
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

func (r *TransactionRepository) PushChanges(ctx context.Context, profileId string, lastPulledAt time.Time, transactionsSync *domain.SyncPayload[*domain.Transaction]) error {
	executor := r.getDB(ctx)
	q := db.New(executor)

	for _, transaction := range transactionsSync.Created {
		params := r.converter.ToDBCreateTransactionParams(transaction)
		_, err := q.CreateTransaction(ctx, params)
		if err != nil {
			slog.Error("failed to store transaction", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store transaction"))
		}
	}

	for _, transaction := range transactionsSync.Updated {
		params := r.converter.ToDBUpdateTransactionParams(transaction)
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

func (r *BaseSyncRepo[DBEntity, DomainEntity]) getDB(ctx context.Context) db.DBTX {
	if tx, ok := extractTx(ctx); ok {
		return tx
	}
	return r.pool
}
