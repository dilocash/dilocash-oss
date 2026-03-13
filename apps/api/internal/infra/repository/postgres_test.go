// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	mappers "github.com/dilocash/dilocash-oss/apps/api/internal/generated/mappers"
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

	t.Run("pull commands changes for empty user", func(t *testing.T) {
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userClean := testUsers[0]

			userID := userClean.ID
			now := time.Now()
			changes, err := commandsRepo.PullChanges(ctx, userID.String(), &now)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			slog.Info("commands received", "changes", changes)
			assert.Len(t, changes.Created, 0)
			assert.Len(t, changes.Updated, 0)
			assert.Len(t, changes.Deleted, 0)
		})
	})

	t.Run("pull commands changes for empty user for the first time", func(t *testing.T) {
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userClean := testUsers[0]

			userID := userClean.ID
			changes, err := commandsRepo.PullChanges(ctx, userID.String(), nil)
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
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithoutData := testUsers[0]

			nilTime := (*time.Time)(nil) // nil/0 time

			timeStr := "2000-01-01T00:00:00.111Z"
			createdAt, err := time.Parse(DateFormatLayout, timeStr)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			var command = &domain.Command{
				ID:            uuid.New(),
				ProfileID:     userWithoutData.ID,
				CommandStatus: 1,
				CreatedAt:     createdAt,
			}
			err = seedCreatedDatabaseCommand(ctx, tx, command)
			if err != nil {
				slog.Error("error seeding created commands", "err", err)
				t.Errorf("expected no error, got %v", err)
			}
			changes, err := commandsRepo.PullChanges(ctx, userWithoutData.ID.String(), nilTime)
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
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithoutData := testUsers[0]

			nilTime := (*time.Time)(nil) // nil/0 time

			timeStr := "2000-01-01T00:00:00.111Z"
			createdAt, err := time.Parse(DateFormatLayout, timeStr)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			var commandId1, commandId2, commandId3 = uuid.New(), uuid.New(), uuid.New()

			var command1 = &domain.Command{
				ID:            commandId1,
				ProfileID:     userWithoutData.ID,
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
				ProfileID:     userWithoutData.ID,
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
				ProfileID:     userWithoutData.ID,
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
			commandsRepo.PushChanges(ctx, userWithoutData.ID.String(), time.Now(), &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{},
				Updated: []*domain.Command{command2},
				Deleted: []uuid.UUID{commandId3},
			})
			changes, err := commandsRepo.PullChanges(ctx, userWithoutData.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			assert.Len(t, changes.Created, 1)
			assert.Len(t, changes.Updated, 1)
			assert.Len(t, changes.Deleted, 1)
		})
	})

	// watermelondb spec: created/updated/deleted record IDs MUST NOT be duplicated
	t.Run("pull commands changes for user with no repeated ids", func(t *testing.T) {
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithData := testUsers[1]

			nilTime := (*time.Time)(nil) // nil/0 time
			changes, err := commandsRepo.PullChanges(ctx, userWithData.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			assert.Len(t, changes.Created, 31)
			assert.Len(t, changes.Updated, 31)
			assert.Len(t, changes.Deleted, 0)

			var returnedIds = make([]uuid.UUID, len(changes.Created))
			for i, command := range changes.Created {
				returnedIds[i] = command.ID
			}
			for _, command := range changes.Updated {
				assert.False(t, slices.Contains(returnedIds, command.ID), "updated command id %v already exists in created commands", command.ID)
			}
		})
	})
}
