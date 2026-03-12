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

-- name: UpdateIntent :one
UPDATE intents
SET text_message = $2, audio_message = $3, image_message = $4, intent_status = $5, requires_review = $6, updated_at = NOW()
WHERE id = $1 AND deleted = false
RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions
SET amount = $2, currency = $3, category = $4, description = $5, updated_at = NOW()
WHERE id = $1 AND deleted = false
RETURNING *;

-- name: DeleteCommand :exec
-- deletes a command and all its related intents and transactions
UPDATE commands
SET deleted = true, updated_at = NOW()
WHERE id = $1 AND deleted = false;
UPDATE transactions
SET deleted = true, updated_at = NOW()
WHERE command_id = $1 AND deleted = false;
UPDATE intents
SET deleted = true, updated_at = NOW()
WHERE command_id = $1 AND deleted = false;

-- name: DeleteTransaction :exec
UPDATE transactions
SET deleted = true, updated_at = NOW()
WHERE id = $1 AND deleted = false;

-- name: DeleteIntent :exec
UPDATE intents
SET deleted = true, updated_at = NOW()
WHERE id = $1 AND deleted = false;

-- name: GetCommandsSync :many
SELECT c.*, 
    CASE 
        WHEN deleted = true THEN 'deleted'
        WHEN created_at > $2 AND (updated_at = c.created_at OR c.updated_at IS NULL) THEN 'created'
        ELSE 'updated'
    END as sync_type
FROM commands c
WHERE c.profile_id = $1 AND c.updated_at > $2
LIMIT $3 OFFSET $4;

-- name: GetIntentsSync :many
SELECT i.*, 
    CASE 
        WHEN deleted = true THEN 'deleted'
        WHEN i.created_at > $2 AND (i.updated_at = i.created_at OR i.updated_at IS NULL) THEN 'created'
        ELSE 'updated'
    END as sync_type
FROM intents i
WHERE i.command_id IN (select id from commands where profile_id = $1) AND i.updated_at > $2
LIMIT $3 OFFSET $4;

-- name: GetTransactionsSync :many
SELECT t.*, 
    CASE 
        WHEN deleted = true THEN 'deleted'
        WHEN t.created_at > $2 AND (t.updated_at = t.created_at OR t.updated_at IS NULL) THEN 'created'
        ELSE 'updated'
    END as sync_type
FROM transactions t
WHERE t.command_id IN (select id from commands where profile_id = $1) AND t.updated_at > $2
LIMIT $3 OFFSET $4;