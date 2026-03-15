// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package repository

import (
	"context"
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
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var pool *pgxpool.Pool // Global variable to hold the database connection

func TestMain(m *testing.M) {
	ctx := context.Background()
	var container *postgres.PostgresContainer
	var err error
	pool, container, err = setupTestPoolDB(ctx)
	if err != nil {
		slog.Error("error setting up test pool db", "err", err)
		os.Exit(1)
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
			now := time.Now()
			commandsRepo.PushChanges(ctx, userWithoutData.ID.String(), &now, &domain.SyncPayload[*domain.Command]{
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

	t.Run("pull commands changes - only created if same updated_at and created_at", func(t *testing.T) {
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithoutData := testUsers[0]
			createdAt := time.Now()
			created := domain.Command{
				ID:            uuid.New(),
				ProfileID:     userWithoutData.ID,
				CommandStatus: 1,
				CreatedAt:     createdAt,
			}
			err = seedCreatedDatabaseCommand(ctx, tx, &created)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			nilTime := (*time.Time)(nil) // nil/0 time
			changes, err := commandsRepo.PullChanges(ctx, userWithoutData.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			assert.Len(t, changes.Created, 1)
			assert.Len(t, changes.Updated, 0)
			assert.Len(t, changes.Deleted, 0)

			assert.Equal(t, created.CreatedAt.Round(time.Millisecond), changes.Created[0].CreatedAt.Round(time.Millisecond))
			assert.Equal(t, created.CreatedAt.Round(time.Millisecond), changes.Created[0].UpdatedAt.Round(time.Millisecond))
			assert.Equal(t, changes.Created[0].CreatedAt.Round(time.Millisecond), changes.Created[0].UpdatedAt.Round(time.Millisecond))
		})
	})
}

func TestPostgresRepo_Commands_PushChanges(t *testing.T) {
	commandsRepo := NewCommandRepository(pool, &mappers.ConverterImpl{})

	t.Run("push commands changes created for empty user", func(t *testing.T) {
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userClean := testUsers[0]
			now := time.Now()

			created := &domain.Command{
				ID:            uuid.New(),
				ProfileID:     userClean.ID,
				CommandStatus: 1,
				CreatedAt:     now,
				UpdatedAt:     now,
				Deleted:       false,
			}

			err = commandsRepo.PushChanges(ctx, userClean.ID.String(), &now, &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{created},
				Updated: []*domain.Command{},
				Deleted: []uuid.UUID{},
			})
			if err != nil {
				slog.Error("error pushing changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			nilTime := (*time.Time)(nil) // nil/0 time
			changes, err := commandsRepo.PullChanges(ctx, userClean.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			assert.Len(t, changes.Created, 1)
			assert.Len(t, changes.Updated, 0)
			assert.Len(t, changes.Deleted, 0)
		})
	})

	// watermelondb spec: deleted rows should not return error if they already do not exist
	t.Run("delete command that do not exist", func(t *testing.T) {
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithoutData := testUsers[0]

			//unexisting id
			unexistingId := uuid.New()

			now := time.Now()
			err = commandsRepo.PushChanges(ctx, userWithoutData.ID.String(), &now, &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{},
				Updated: []*domain.Command{},
				Deleted: []uuid.UUID{unexistingId},
			})
			if err != nil {
				slog.Error("error pushing changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			nilTime := (*time.Time)(nil) // nil/0 time
			changes, err := commandsRepo.PullChanges(ctx, userWithoutData.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			assert.Len(t, changes.Created, 0)
			assert.Len(t, changes.Updated, 0)
			assert.Len(t, changes.Deleted, 0)

		})
	})

	t.Run("update command for existing item", func(t *testing.T) {
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

			toUpdate := changes.Created[0]
			toUpdateId := toUpdate.ID

			toUpdate.CommandStatus = 4
			newUpdatedTime := time.Now()
			toUpdate.UpdatedAt = newUpdatedTime

			now := time.Now()
			err = commandsRepo.PushChanges(ctx, userWithData.ID.String(), &now, &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{},
				Updated: []*domain.Command{toUpdate},
				Deleted: []uuid.UUID{},
			})
			if err != nil {
				slog.Error("error pushing changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}

			changes, err = commandsRepo.PullChanges(ctx, userWithData.ID.String(), nilTime)
			if err != nil {
				slog.Error("error pulling changes", "err", err)
				t.Errorf("expected no error, got %v", err)
			}
			for _, command := range changes.Updated {
				if command.ID == toUpdateId {
					// 100 ms allowed difference
					assert.GreaterOrEqual(t, command.UpdatedAt.Round(time.Millisecond), newUpdatedTime.Round(time.Millisecond).Add(-time.Millisecond*100))
					assert.LessOrEqual(t, command.UpdatedAt.Round(time.Millisecond), newUpdatedTime.Round(time.Millisecond).Add(time.Millisecond*100))
					assert.Equal(t, command.CommandStatus, toUpdate.CommandStatus)
					break
				}
			}
			assert.Len(t, changes.Created, 30)
			assert.Len(t, changes.Updated, 32)
			assert.Len(t, changes.Deleted, 0)

		})
	})

	// watermelondb spec: return an error code (to force frontend to pull the information about this deleted ID
	t.Run("update command that was deleted and then updated", func(t *testing.T) {
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
			toDelete := changes.Created[0]
			deletedId := toDelete.ID

			now := time.Now()
			err = commandsRepo.PushChanges(ctx, userWithData.ID.String(), &now, &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{},
				Updated: []*domain.Command{},
				Deleted: []uuid.UUID{deletedId},
			})

			updatedAfterDeleted := &domain.Command{
				ID:            deletedId,
				ProfileID:     userWithData.ID,
				CommandStatus: 1,
				CreatedAt:     toDelete.CreatedAt,
				UpdatedAt:     toDelete.UpdatedAt,
				Deleted:       false,
			}
			err = commandsRepo.PushChanges(ctx, userWithData.ID.String(), &now, &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{},
				Updated: []*domain.Command{updatedAfterDeleted},
				Deleted: []uuid.UUID{},
			})

			if err == nil {
				t.Errorf("expected an error, but got nil")
			}
		})
	})

	// watermelondb spec: you MUST create it, and MUST NOT return an error code
	t.Run("update command that never existed", func(t *testing.T) {
		testWithRollback(t, pool, func(ctx context.Context, tx pgx.Tx) {
			// add test users inside transaction
			testUsers, err := seedDatabaseUsers(ctx, tx)
			if err != nil {
				t.Fatal(err)
			}
			userWithoutData := testUsers[0]

			nilTime := (*time.Time)(nil) // nil/0 time
			now := time.Now()

			//unexisting id
			unexistingId := uuid.New()

			updatedButNeverExisted := &domain.Command{
				ID:            unexistingId,
				ProfileID:     userWithoutData.ID,
				CommandStatus: 1,
				CreatedAt:     now,
				UpdatedAt:     now,
				Deleted:       false,
			}
			err = commandsRepo.PushChanges(ctx, userWithoutData.ID.String(), &now, &domain.SyncPayload[*domain.Command]{
				Created: []*domain.Command{},
				Updated: []*domain.Command{updatedButNeverExisted},
				Deleted: []uuid.UUID{},
			})

			assert.NoError(t, err)

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
}
