-- Copyright (c) 2026 dilocash
-- Use of this source code is governed by an MIT-style
-- license that can be found in the LICENSE file.

-- name: GetProfile :one
SELECT * FROM profiles
WHERE id = $1 LIMIT 1;

-- name: CreateCommand :one
INSERT INTO commands (
    id, profile_id, command_status, created_at
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: CreateTransaction :one
INSERT INTO transactions (
    id, command_id, amount, currency, category, description, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: CreateIntent :one
INSERT INTO intents (
    id, command_id, text_message, audio_message, image_message, intent_status, requires_review, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateCommand :one
UPDATE commands
SET command_status = $2, updated_at = NOW()
WHERE id = $1 AND deleted = false
RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions
SET amount = $2, currency = $3, category = $4, description = $5, updated_at = NOW()
WHERE id = $1 AND deleted = false
RETURNING *;

-- name: UpdateIntent :one
UPDATE intents
SET text_message = $2, audio_message = $3, image_message = $4, intent_status = $5, requires_review = $6, updated_at = NOW()
WHERE id = $1 AND deleted = false
RETURNING *;

-- name: DeleteCommand :exec
UPDATE commands
SET deleted = true, updated_at = NOW()
WHERE id = $1 AND deleted = false;

-- name: DeleteTransaction :exec
UPDATE transactions
SET deleted = true, updated_at = NOW()
WHERE id = $1 AND deleted = false;

-- name: DeleteIntent :exec
UPDATE intents
SET deleted = true, updated_at = NOW()
WHERE id = $1 AND deleted = false;

-- name: ListCommandsByProfileId :many
-- initial sync
SELECT c.* FROM commands c
WHERE c.profile_id = $1 AND c.deleted = false
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListTransactionsByProfileId :many
-- initial sync
SELECT t.* FROM transactions t, commands c
WHERE c.id = t.command_id AND c.profile_id = $1 AND t.deleted = false
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListIntentsByProfileId :many
-- initial sync
SELECT i.* FROM intents i, commands c
WHERE c.id = i.command_id AND c.profile_id = $1 AND i.deleted = false
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListCommandsByProfileIdAndCreatedAfter :many
SELECT c.*
FROM commands c
WHERE c.profile_id = $1
  AND c.deleted = false
  AND c.created_at > $2 and c.updated_at <= $2 -- only new records
ORDER BY c.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListIntentsByProfileIdAndCreatedAfter :many
SELECT i.*
FROM intents i
WHERE i.command_id IN (select id from commands where profile_id = $1)
  AND i.deleted = false
  AND i.created_at > $2 and i.updated_at <= $2 -- only new records
ORDER BY i.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListTransactionsByProfileIdAndCreatedAfter :many
SELECT t.*
FROM transactions t
WHERE t.command_id IN (select id from commands where profile_id = $1)
  AND t.deleted = false
  AND t.created_at > $2 and t.updated_at <= $2 -- only new records
ORDER BY t.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListCommandsByProfileIdAndUpdatedAfter :many
SELECT c.*
FROM commands c
WHERE c.profile_id = $1
  AND c.updated_at > c.created_at AND c.updated_at > $2 AND c.deleted = false
ORDER BY c.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListIntentsByProfileIdAndUpdatedAfter :many
SELECT i.*
FROM intents i
WHERE i.command_id IN (select id from commands where profile_id = $1)
  AND i.updated_at > i.created_at AND i.updated_at > $2 AND i.deleted = false
ORDER BY i.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListTransactionsByProfileIdAndUpdatedAfter :many
SELECT t.*
FROM transactions t
WHERE t.command_id IN (select id from commands where profile_id = $1)
  AND t.updated_at > t.created_at AND t.updated_at > $2 AND t.deleted = false
ORDER BY t.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListDeletedCommandsByProfileIdAndUpdatedAfter :many
SELECT c.id as command_id
FROM commands c
WHERE c.profile_id = $1
  AND c.updated_at > c.created_at AND c.updated_at > $2 AND c.deleted = true
ORDER BY c.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListDeletedIntentsByProfileIdAndUpdatedAfter :many
SELECT i.id as intent_id
FROM intents i
WHERE i.command_id IN (select id from commands where profile_id = $1)
  AND i.updated_at > i.created_at AND i.updated_at > $2 AND i.deleted = true
ORDER BY i.created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListDeletedTransactionsByProfileIdAndUpdatedAfter :many
SELECT t.id as transaction_id
FROM transactions t
WHERE t.command_id IN (select id from commands where profile_id = $1)
  AND t.updated_at > t.created_at AND t.updated_at > $2 AND t.deleted = true
ORDER BY t.created_at DESC
LIMIT $3 OFFSET $4;