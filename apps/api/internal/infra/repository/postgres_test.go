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

	mappers "github.com/dilocash/dilocash-oss/apps/api/internal/generated/mappers"
	"github.com/friendliai/atlas-go-sdk/atlasexec"
	"github.com/google/uuid"
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
	ctx := context.Background()

	// add test users
	testUsers, err := seedDatabaseUsers(ctx, pool)
	if err != nil {
		t.Fatal(err)
	}

	var userClean = testUsers[0]
	var userWithData = testUsers[1]

	// instance the repository
	commandsRepo := NewCommandRepository(pool, &mappers.ConverterImpl{})

	// run the test
	t.Run("pull commands changes for first user", func(t *testing.T) {
		userID := userClean.ID
		changes, err := commandsRepo.PullChanges(ctx, userID.String(), time.Now())
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		slog.Info("commands received", "changes", changes)
		assert.Len(t, changes.Created, 0)
		assert.Len(t, changes.Deleted, 0)
		assert.Len(t, changes.Updated, 0)
	})

	t.Run("pull commands changes for user with only 1 command created", func(t *testing.T) {

		nilTime := time.Time{} // nil/0 time

		layout := "2006-01-02T15:04:05.999Z"
		timeStr := "2000-01-01T00:00:00.111Z"
		createdAt, err := time.Parse(layout, timeStr)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		err = seedCreatedDatabaseCommands(ctx, pool, uuid.New(), userWithData.ID, 1, createdAt)
		if err != nil {
			slog.Error("error seeding created commands", "err", err)
			t.Errorf("expected no error, got %v", err)
		}
		changes, err := commandsRepo.PullChanges(ctx, userWithData.ID.String(), nilTime)
		if err != nil {
			slog.Error("error pulling changes", "err", err)
			t.Errorf("expected no error, got %v", err)
		}
		_ = changes

		assert.Len(t, changes.Created, 1)
		assert.Len(t, changes.Deleted, 0)
		assert.Len(t, changes.Updated, 0)
	})
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

func seedDatabaseUsers(ctx context.Context, pool *pgxpool.Pool) ([]TestUser, error) {
	users := []TestUser{
		{ID: uuid.New(), Email: "empty@mail.com"},
		{ID: uuid.New(), Email: "2026-jan-daily-data@mail.com"},
	}

	for _, u := range users {
		seedSQL := `INSERT INTO auth.users (id, email, aud, role) VALUES ($1, $2, 'authenticated', 'authenticated');`
		_, err := pool.Exec(ctx, seedSQL, u.ID, u.Email)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func seedCreatedDatabaseCommands(ctx context.Context, pool *pgxpool.Pool, id uuid.UUID, profileID uuid.UUID, commandStatus int, createdAt time.Time) error {
	slog.Info("seed created command", "id", id, "profileID", profileID, "commandStatus", commandStatus, "createdAt", createdAt)

	seedSQL := `
		INSERT INTO commands (id, profile_id, command_status, created_at, updated_at, deleted) 
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := pool.Exec(ctx, seedSQL, id, profileID.String(), commandStatus, createdAt, createdAt, false)
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
