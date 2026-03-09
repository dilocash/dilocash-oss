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
	ctx  *context.Context
	q    *db.Queries
	pool *pgxpool.Pool
}

func NewPostgresRepo(ctx *context.Context, pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		ctx:  ctx,
		q:    db.New(pool),
		pool: pool,
	}
}

func (r *PostgresRepo) PullCommandChanges(profileId string, lastPulledAt time.Time) (*domain.CommandsSync, error) {
	converter := &mappers.ConverterImpl{}
	created := []domain.Command{}

	slog.Info("pulling command changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	// execute query to get created commands since lastPulledAt
	createdCommandsResult, err := r.q.ListCommandsByProfileIdAndCreatedAfter(*r.ctx, db.ListCommandsByProfileIdAndCreatedAfterParams{
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
	updatedCommandsResult, err := r.q.ListCommandsByProfileIdAndUpdatedAfter(*r.ctx, db.ListCommandsByProfileIdAndUpdatedAfterParams{
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
	deletedCommandsResult, err := r.q.ListDeletedCommandsByProfileIdAndUpdatedAfter(*r.ctx, db.ListDeletedCommandsByProfileIdAndUpdatedAfterParams{
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

func (r *PostgresRepo) PullIntentChanges(profileId string, lastPulledAt time.Time) (*domain.IntentsSync, error) {
	converter := &mappers.ConverterImpl{}
	created := []domain.Intent{}

	slog.Info("pulling intent changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	// execute query to get intents since lastPulledAt
	createdIntentsResult, err := r.q.ListIntentsByProfileIdAndCreatedAfter(*r.ctx, db.ListIntentsByProfileIdAndCreatedAfterParams{
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
	updatedIntentsResult, err := r.q.ListIntentsByProfileIdAndUpdatedAfter(*r.ctx, db.ListIntentsByProfileIdAndUpdatedAfterParams{
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
	deletedIntentsResult, err := r.q.ListDeletedIntentsByProfileIdAndUpdatedAfter(*r.ctx, db.ListDeletedIntentsByProfileIdAndUpdatedAfterParams{
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

func (r *PostgresRepo) PullTransactionChanges(profileId string, lastPulledAt time.Time) (*domain.TransactionsSync, error) {
	converter := &mappers.ConverterImpl{}
	created := []domain.Transaction{}

	slog.Info("pulling transaction changes", "profileId", profileId, "lastPulledAt", lastPulledAt)
	// execute query to get transactions since lastPulledAt
	createdTransactionsResult, err := r.q.ListTransactionsByProfileIdAndCreatedAfter(*r.ctx, db.ListTransactionsByProfileIdAndCreatedAfterParams{
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
	updatedTransactionsResult, err := r.q.ListTransactionsByProfileIdAndUpdatedAfter(*r.ctx, db.ListTransactionsByProfileIdAndUpdatedAfterParams{
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
	deletedTransactionsResult, err := r.q.ListDeletedTransactionsByProfileIdAndUpdatedAfter(*r.ctx, db.ListDeletedTransactionsByProfileIdAndUpdatedAfterParams{
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

func (r *PostgresRepo) PushCommandChanges(profileId string, commandsSync *domain.CommandsSync) error {
	tx, err := r.pool.Begin(*r.ctx)
	if err != nil {
		return connect.NewError(connect.CodeInternal, errors.New("failed to begin transaction"))
	}
	defer tx.Rollback(*r.ctx)

	qtx := r.q.WithTx(tx)

	for _, command := range commandsSync.Created {
		params := db.CreateCommandParams{
			ID:            command.ID,
			ProfileID:     uuid.MustParse(profileId),
			CommandStatus: command.CommandStatus,
			CreatedAt:     command.CreatedAt,
		}
		_, err := qtx.CreateCommand(*r.ctx, params)
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
		_, err := qtx.UpdateCommand(*r.ctx, params)
		if err != nil {
			slog.Error("failed to store command", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store command"))
		}
	}

	for _, command := range commandsSync.Deleted {
		err := qtx.DeleteCommand(*r.ctx, command)
		if err != nil {
			slog.Error("failed to delete command", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to delete command"))
		}
	}
	res := tx.Commit(*r.ctx)
	if res != nil {
		slog.Error("failed to commit transaction", "error", res)
		return connect.NewError(connect.CodeInternal, errors.New("failed to commit transaction"))
	}
	return nil
}

func (r *PostgresRepo) PushIntentChanges(profileId string, intentsSync *domain.IntentsSync) error {
	tx, err := r.pool.Begin(*r.ctx)
	if err != nil {
		return connect.NewError(connect.CodeInternal, errors.New("failed to begin transaction"))
	}
	defer tx.Rollback(*r.ctx)

	qtx := r.q.WithTx(tx)

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
		_, err := qtx.CreateIntent(*r.ctx, params)
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
		_, err := qtx.UpdateIntent(*r.ctx, params)
		if err != nil {
			slog.Error("failed to store intent", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store intent"))
		}
	}

	for _, intent := range intentsSync.Deleted {
		err := qtx.DeleteIntent(*r.ctx, intent)
		if err != nil {
			slog.Error("failed to delete intent", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to delete intent"))
		}
	}
	res := tx.Commit(*r.ctx)
	if res != nil {
		slog.Error("failed to commit transaction", "error", res)
		return connect.NewError(connect.CodeInternal, errors.New("failed to commit transaction"))
	}
	return nil
}

func (r *PostgresRepo) PushTransactionChanges(profileId string, transactionsSync *domain.TransactionsSync) error {
	tx, err := r.pool.Begin(*r.ctx)
	if err != nil {
		return connect.NewError(connect.CodeInternal, errors.New("failed to begin transaction"))
	}
	defer tx.Rollback(*r.ctx)

	qtx := r.q.WithTx(tx)

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
		_, err := qtx.CreateTransaction(*r.ctx, params)
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
		_, err := qtx.UpdateTransaction(*r.ctx, params)
		if err != nil {
			slog.Error("failed to store transaction", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to store transaction"))
		}
	}

	for _, intent := range transactionsSync.Deleted {
		err := qtx.DeleteTransaction(*r.ctx, intent)
		if err != nil {
			slog.Error("failed to delete transaction", "error", err)
			return connect.NewError(connect.CodeInternal, errors.New("failed to delete transaction"))
		}
	}
	res := tx.Commit(*r.ctx)
	if res != nil {
		slog.Error("failed to commit transaction", "error", res)
		return connect.NewError(connect.CodeInternal, errors.New("failed to commit transaction"))
	}
	return nil
}
