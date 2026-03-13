package repository

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	db "github.com/dilocash/dilocash-oss/apps/api/internal/generated/db/postgres"
	"github.com/friendliai/atlas-go-sdk/atlasexec"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const DateFormatLayout string = "2006-01-02T15:04:05.999Z" // date formatting layout

type TestUser struct {
	ID    uuid.UUID
	Email string
}

func seedDatabaseUsers(ctx context.Context, db db.DBTX) ([]TestUser, error) {
	users := []TestUser{
		{ID: uuid.New(), Email: "empty@mail.com"},
		{ID: uuid.New(), Email: "2026-jan-daily-data@mail.com"},
	}

	for _, u := range users {
		seedSQL := `INSERT INTO auth.users (id, email, aud, role) VALUES ($1, $2, 'authenticated', 'authenticated');`
		_, err := db.Exec(ctx, seedSQL, u.ID, u.Email)
		if err != nil {
			return nil, err
		}
	}
	dates := []string{}
	// dates for 2026-01-01 month
	for i := 1; i <= 31; i++ {
		dates = append(dates, fmt.Sprintf("2026-01-%02dT00:00:00.000Z", i))
	}
	userWithData := users[1]
	for _, dateStr := range dates {
		createdAt, err := time.Parse(DateFormatLayout, dateStr)
		if err != nil {
			return nil, err
		}
		created := domain.Command{
			ID:            uuid.New(),
			ProfileID:     userWithData.ID,
			CommandStatus: 1,
			CreatedAt:     createdAt,
		}
		err = seedCreatedDatabaseCommand(ctx, db, &created)
		if err != nil {
			return nil, err
		}
		updatedAt := createdAt.Add(time.Hour)
		updated := domain.Command{
			ID:            uuid.New(),
			ProfileID:     userWithData.ID,
			CommandStatus: 1,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		}
		err = seedUpdatedDatabaseCommand(ctx, db, &updated)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func seedCreatedDatabaseCommand(ctx context.Context, db db.DBTX, command *domain.Command) error {
	slog.Debug("seed created command", "id", command.ID, "profileID", command.ProfileID, "commandStatus", command.CommandStatus, "createdAt", command.CreatedAt)

	seedSQL := `
		INSERT INTO commands (id, profile_id, command_status, created_at, updated_at, deleted) 
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := db.Exec(ctx, seedSQL, command.ID.String(), command.ProfileID.String(), command.CommandStatus, command.CreatedAt, command.CreatedAt, false)
	if err != nil {
		return err
	}

	return nil
}

func seedUpdatedDatabaseCommand(ctx context.Context, db db.DBTX, command *domain.Command) error {
	slog.Debug("seed updated command", "id", command.ID, "profileID", command.ProfileID, "commandStatus", command.CommandStatus, "createdAt", command.CreatedAt, "updatedAt", command.UpdatedAt)

	seedSQL := `
		INSERT INTO commands (id, profile_id, command_status, created_at, updated_at, deleted) 
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := db.Exec(ctx, seedSQL, command.ID.String(), command.ProfileID.String(), command.CommandStatus, command.CreatedAt, command.UpdatedAt, false)
	if err != nil {
		return err
	}

	return nil
}

func seedCreatedDatabaseIntent(ctx context.Context, db db.DBTX, intent *domain.Intent) error {
	slog.Debug("seed created intent", "id", intent.ID, "commandID", intent.CommandID, "intentStatus", intent.IntentStatus, "createdAt", intent.CreatedAt)

	seedSQL := `
		INSERT INTO intents (id, command_id, text_message, audio_message, image_message, intent_status, requires_review, created_at, updated_at, deleted) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`
	_, err := db.Exec(ctx, seedSQL, intent.ID.String(), intent.CommandID.String(), intent.TextMessage, intent.AudioMessage, intent.ImageMessage, intent.IntentStatus, intent.RequiresReview, intent.CreatedAt, intent.CreatedAt, intent.Deleted)
	if err != nil {
		return err
	}

	return nil
}

func seedCreatedDatabaseTransaction(ctx context.Context, db db.DBTX, transaction *domain.Transaction) error {
	slog.Debug("seed created transaction", "id", transaction.ID, "commandID", transaction.CommandID, "amount", transaction.Amount, "currency", transaction.Currency, "category", transaction.Category, "description", transaction.Description, "createdAt", transaction.CreatedAt)

	seedSQL := `
		INSERT INTO transactions (id, command_id, amount, currency, category, description, created_at, updated_at, deleted) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`
	_, err := db.Exec(ctx, seedSQL, transaction.ID.String(), transaction.CommandID.String(), transaction.Amount, transaction.Currency, transaction.Category, transaction.Description, transaction.CreatedAt, transaction.CreatedAt, transaction.Deleted)
	if err != nil {
		return err
	}

	return nil
}

func setupSupabaseSchema(ctx context.Context, pool *pgxpool.Pool) error {
	authSchema, _ := os.ReadFile("../testing/supabase_auth_schema.sql")
	if _, err := pool.Exec(ctx, string(authSchema)); err != nil {
		return err
	}
	return nil
}

// testWithRollback wraps a test in a transaction and rolls it back
func testWithRollback(t *testing.T, pool *pgxpool.Pool, testFunc func(ctx context.Context, tx pgx.Tx)) {
	t.Helper()
	ctx := context.Background()

	// Begin a transaction for this specific test
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	// Use injectTx from tx_context.go instead of manual context manipulation
	ctx = injectTx(ctx, tx)

	// Roll back the transaction at the end of the test
	// This ensures a clean state for the next test
	t.Cleanup(func() {
		if err := tx.Rollback(ctx); err != nil {
			// Rollback might fail if the transaction was already committed or closed,
			// but in tests we generally expect it to succeed unless we committed on purpose.
			slog.Debug("rollback finallized", "err", err)
		}
	})

	// Run the actual test logic with the transaction
	testFunc(ctx, tx)
}

func setupMigrations(ctx context.Context, connStr string) error {
	slog.Info("start migrations", "connStr", connStr)
	// use atlas sdk to apply migrations
	client, err := atlasexec.NewClient(".", "atlas")
	if err != nil {
		slog.Error("Error creating atlas client", slog.Any("err", err))
		return err
	}

	// apply migrations that are in internal/infrastructure/db/migrations
	_, err = client.MigrateApply(ctx, &atlasexec.MigrateApplyParams{
		URL:        connStr, // container url
		DirURL:     "file://../../../migrations",
		AllowDirty: true,
	})
	if err != nil {
		slog.Error("Error applying migrations", slog.Any("err", err))
		return err
	}
	return nil
}
