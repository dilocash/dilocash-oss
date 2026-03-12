// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

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
	mappers "github.com/dilocash/dilocash-oss/apps/api/internal/generated/mappers"
	"github.com/friendliai/atlas-go-sdk/atlasexec"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var pool *pgxpool.Pool // Global variable to hold the database connection

func TestMain(m *testing.M) {
	ctx := context.Background()
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
	container, err := postgres.Run(ctx,
		"postgres:17.9-alpine", // Specify the Docker image
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
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

	pool, err = pgxpool.New(ctx, connStr)
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

	defer pool.Close()

	// Run all tests in the package
	code := m.Run()

	// Teardown logic (e.g., cleaning up resources)
	slog.Info("Performing teardown...")
	// ... cleanup code ...
	// Teardown: terminate the container after all tests are done
	if err = container.Terminate(ctx); err != nil {
		slog.Error("failed to terminate container", slog.Any("err", err))
	}

	// Exit with the appropriate code
	os.Exit(code)
}

func TestPostgresRepo_Commands_PullChanges(t *testing.T) {
	commandsRepo := NewCommandRepository(pool, &mappers.ConverterImpl{})

	t.Run("pull commands changes for first user", func(t *testing.T) {
		testWithRollback(t, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userClean := testUsers[0]

			userID := userClean.ID
			changes, err := commandsRepo.PullChanges(ctx, userID.String(), time.Now())
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			slog.Info("commands received", "changes", changes)
			assert.Len(t, changes.Created, 0)
			assert.Len(t, changes.Updated, 0)
			assert.Len(t, changes.Deleted, 0)
		})
	})

	t.Run("pull commands changes for user with only 1 command created", func(t *testing.T) {
		testWithRollback(t, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithData := testUsers[1]

			nilTime := time.Time{} // nil/0 time

			layout := "2006-01-02T15:04:05.999Z"
			timeStr := "2000-01-01T00:00:00.111Z"
			createdAt, err := time.Parse(layout, timeStr)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			var command = &domain.Command{
				ID:            uuid.New(),
				ProfileID:     userWithData.ID,
				CommandStatus: 1,
				CreatedAt:     createdAt,
			}
			err = seedCreatedDatabaseCommand(ctx, tx, command)
			if err != nil {
				slog.Error("error seeding created commands", "err", err)
				t.Errorf("expected no error, got %v", err)
			}
			changes, err := commandsRepo.PullChanges(ctx, userWithData.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			assert.Len(t, changes.Created, 1)
			assert.Len(t, changes.Updated, 0)
			assert.Len(t, changes.Deleted, 0)
		})
	})

	t.Run("pull commands changes for user with 1 command created, 1 updated and 1 deleted", func(t *testing.T) {
		testWithRollback(t, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithData := testUsers[1]

			nilTime := time.Time{} // nil/0 time

			layout := "2006-01-02T15:04:05.999Z"
			timeStr := "2000-01-01T00:00:00.111Z"
			createdAt, err := time.Parse(layout, timeStr)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			var commandId1, commandId2, commandId3 = uuid.New(), uuid.New(), uuid.New()

			var command1 = &domain.Command{
				ID:            commandId1,
				ProfileID:     userWithData.ID,
				CommandStatus: 1,
				CreatedAt:     createdAt,
				Deleted:       false,
			}
			err = seedCreatedDatabaseCommand(ctx, tx, command1)
			if err != nil {
				slog.Error("error seeding created commands", "err", err)
				t.Errorf("expected no error, got %v", err)
			}
			command2 := &domain.Command{
				ID:            commandId2,
				ProfileID:     userWithData.ID,
				CommandStatus: 1,
				CreatedAt:     createdAt,
				UpdatedAt:     createdAt,
				Deleted:       false,
			}
			err = seedUpdatedDatabaseCommand(ctx, tx, command2)
			if err != nil {
				slog.Error("error seeding updated commands", "err", err)
				t.Errorf("expected no error, got %v", err)
			}
			command3 := &domain.Command{
				ID:            commandId3,
				ProfileID:     userWithData.ID,
				CommandStatus: 1,
				CreatedAt:     createdAt,
				UpdatedAt:     createdAt,
				Deleted:       true,
			}
			err = seedUpdatedDatabaseCommand(ctx, tx, command3)
			if err != nil {
				slog.Error("error seeding updated commands", "err", err)
				t.Errorf("expected no error, got %v", err)
			}
			commandsRepo.PushChanges(ctx, userWithData.ID.String(), time.Now(), &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{},
				Updated: []*domain.Command{command2},
				Deleted: []uuid.UUID{commandId3},
			})
			changes, err := commandsRepo.PullChanges(ctx, userWithData.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			assert.Len(t, changes.Created, 1)
			assert.Len(t, changes.Updated, 1)
			assert.Len(t, changes.Deleted, 1)
		})
	})
}

// testWithRollback wraps a test in a transaction and rolls it back
func testWithRollback(t *testing.T, testFunc func(ctx context.Context, tx pgx.Tx)) {
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
	return users, nil
}

func seedCreatedDatabaseCommand(ctx context.Context, db db.DBTX, command *domain.Command) error {
	slog.Info("seed created command", "id", command.ID, "profileID", command.ProfileID, "commandStatus", command.CommandStatus, "createdAt", command.CreatedAt)

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
	slog.Info("seed updated command", "id", command.ID, "profileID", command.ProfileID, "commandStatus", command.CommandStatus, "createdAt", command.CreatedAt, "updatedAt", command.UpdatedAt)

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
	slog.Info("seed created intent", "id", intent.ID, "commandID", intent.CommandID, "intentStatus", intent.IntentStatus, "createdAt", intent.CreatedAt)

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
	slog.Info("seed created transaction", "id", transaction.ID, "commandID", transaction.CommandID, "amount", transaction.Amount, "currency", transaction.Currency, "category", transaction.Category, "description", transaction.Description, "createdAt", transaction.CreatedAt)

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
