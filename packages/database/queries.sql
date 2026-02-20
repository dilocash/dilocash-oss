-- Copyright (c) 2026 dilocash
-- Use of this source code is governed by an MIT-style
-- license that can be found in the LICENSE file.

-- name: GetProfile :one
SELECT * FROM profiles
WHERE id = $1 LIMIT 1;

-- name: CreateCommand :one
INSERT INTO commands (
    id, profile_id, command_status
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: CreateTransaction :one
INSERT INTO transactions (
    command_id, amount, currency, category, description
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateIntent :one
INSERT INTO intents (
    command_id, text_message, audio_message, image_message, intent_status, requires_review
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListTransactionsByProfileId :many
SELECT * FROM transactions t, commands c
WHERE c.id = t.command_id AND c.profile_id = $1
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListIntentsByProfileId :many
SELECT * FROM intents i, commands c
WHERE c.id = i.command_id AND c.profile_id = $1
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;


-- name: ListCommandsByProfileId :many
SELECT * FROM commands c
WHERE c.profile_id = $1
AND updated_at > $2
ORDER BY c.created_at DESC
LIMIT $3 OFFSET $4;
