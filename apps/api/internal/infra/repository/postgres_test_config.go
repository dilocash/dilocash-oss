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
	"github.com/shopspring/decimal"
	"github.com/testcontainers/testcontainers-go"
	postgrestestcontainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const DateFormatLayout string = "2006-01-02T15:04:05.999Z" // date formatting layout

type TestUser struct {
	ID                  uuid.UUID
	Email               string
	CreatedCommands     []domain.Command
	CreatedIntents      []domain.Intent
	CreatedTransactions []domain.Transaction
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
	for _, dateStr := range dates {
		createdAt, err := time.Parse(DateFormatLayout, dateStr)
		if err != nil {
			return nil, err
		}
		created := domain.Command{
			ID:            uuid.New(),
			ProfileID:     users[1].ID,
			CommandStatus: 1,
			CreatedAt:     createdAt,
		}
		err = seedCreatedDatabaseCommand(ctx, db, &created)
		if err != nil {
			return nil, err
		}
		users[1].CreatedCommands = append(users[1].CreatedCommands, created)
		createdIntent := domain.Intent{
			ID:           uuid.New(),
			CommandID:    created.ID,
			IntentStatus: 1,
			CreatedAt:    created.CreatedAt,
			UpdatedAt:    created.CreatedAt,
			Deleted:      false,
		}
		err = seedCreatedDatabaseIntent(ctx, db, &createdIntent)
		if err != nil {
			return nil, err
		}
		users[1].CreatedIntents = append(users[1].CreatedIntents, createdIntent)
		createdTransaction := domain.Transaction{
			ID:          uuid.New(),
			CommandID:   created.ID,
			Amount:      decimal.NewFromFloat(100.0),
			Currency:    "USD",
			Category:    "Category",
			Description: "Description",
			CreatedAt:   created.CreatedAt,
			UpdatedAt:   created.CreatedAt,
			Deleted:     false,
		}
		err = seedCreatedDatabaseTransaction(ctx, db, &createdTransaction)
		if err != nil {
			return nil, err
		}
		users[1].CreatedTransactions = append(users[1].CreatedTransactions, createdTransaction)

		updatedAt := createdAt.Add(time.Hour)
		updated := domain.Command{
			ID:            uuid.New(),
			ProfileID:     users[1].ID,
			CommandStatus: 1,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		}
		err = seedUpdatedDatabaseCommand(ctx, db, &updated)
		if err != nil {
			return nil, err
		}
		intent := domain.Intent{
			ID:           uuid.New(),
			CommandID:    updated.ID,
			IntentStatus: 1,
			CreatedAt:    updated.CreatedAt,
			UpdatedAt:    updated.UpdatedAt,
			Deleted:      false,
		}
		err = seedUpdatedDatabaseIntent(ctx, db, &intent)
		if err != nil {
			return nil, err
		}
		transaction := domain.Transaction{
			ID:          uuid.New(),
			CommandID:   updated.ID,
			Amount:      decimal.NewFromFloat(100.0),
			Currency:    "USD",
			Category:    "Category",
			Description: "Description",
			CreatedAt:   updated.CreatedAt,
			UpdatedAt:   updated.UpdatedAt,
			Deleted:     false,
		}
		err = seedUpdatedDatabaseTransaction(ctx, db, &transaction)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func seedCreatedDatabaseCommand(ctx context.Context, db db.DBTX, command *domain.Command) error {
	// slog.Debug("seed created command", "id", command.ID, "profileID", command.ProfileID, "commandStatus", command.CommandStatus, "createdAt", command.CreatedAt)

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
	// slog.Debug("seed updated command", "id", command.ID, "profileID", command.ProfileID, "commandStatus", command.CommandStatus, "createdAt", command.CreatedAt, "updatedAt", command.UpdatedAt)

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
	// slog.Debug("seed created intent", "id", intent.ID, "commandID", intent.CommandID, "intentStatus", intent.IntentStatus, "createdAt", intent.CreatedAt)

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

func seedUpdatedDatabaseIntent(ctx context.Context, db db.DBTX, intent *domain.Intent) error {
	seedSQL := `
		INSERT INTO intents (id, command_id, text_message, audio_message, image_message, intent_status, requires_review, created_at, updated_at, deleted) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`
	_, err := db.Exec(ctx, seedSQL, intent.ID.String(), intent.CommandID.String(), intent.TextMessage, intent.AudioMessage, intent.ImageMessage, intent.IntentStatus, intent.RequiresReview, intent.CreatedAt, intent.UpdatedAt, intent.Deleted)
	if err != nil {
		return err
	}

	return nil
}

func seedUpdatedDatabaseTransaction(ctx context.Context, db db.DBTX, transaction *domain.Transaction) error {
	slog.Debug("seed updated transaction", "id", transaction.ID, "commandID", transaction.CommandID, "amount", transaction.Amount, "currency", transaction.Currency, "category", transaction.Category, "description", transaction.Description, "createdAt", transaction.CreatedAt, "updatedAt", transaction.UpdatedAt)

	seedSQL := `
		INSERT INTO transactions (id, command_id, amount, currency, category, description, created_at, updated_at, deleted) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`
	_, err := db.Exec(ctx, seedSQL, transaction.ID.String(), transaction.CommandID.String(), transaction.Amount, transaction.Currency, transaction.Category, transaction.Description, transaction.CreatedAt, transaction.UpdatedAt, transaction.Deleted)
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

func setupTestPoolDB(ctx context.Context) (*pgxpool.Pool, *postgrestestcontainer.PostgresContainer, error) {
	fmt.Println("Performing setup...")

	logLevel := new(slog.LevelVar)

	// Example of dynamic level setting (simplified)
	var handler slog.Handler
	logLevel.Set(slog.LevelDebug)
	handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	})
	// Set the default logger
	slog.SetDefault(slog.New(handler))
	// Run the PostgreSQL container
	container, err := postgrestestcontainer.Run(ctx,
		"postgres:17.9-alpine", // Specify the Docker image
		postgrestestcontainer.WithDatabase("testdb"),
		postgrestestcontainer.WithUsername("user"),
		postgrestestcontainer.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2), // Wait for the ready signal
		),
	)
	if err != nil {
		slog.Error("Error creating postgres container", slog.Any("err", err))
	}

	// Get the connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		slog.Error("failed to get connection string", slog.Any("err", err))
	}

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		slog.Error("failed to create pool", slog.Any("err", err))
	}
	// Ensure the connection is valid
	if err = pool.Ping(ctx); err != nil {
		slog.Error("failed to ping database", slog.Any("err", err))
	}

	// setup supabase schema
	err = setupSupabaseSchema(ctx, pool)
	if err != nil {
		slog.Error("error seting up supabase schema", slog.Any("err", err))
	}

	// run setupMigrations migrations
	err = setupMigrations(ctx, connStr)
	if err != nil {
		slog.Error("error applying migrations", slog.Any("err", err))
	}

	return pool, container, nil
}
