-- Copyright (c) 2026 dilocash
-- Use of this source code is governed by an MIT-style
-- license that can be found in the LICENSE file.

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    id, email, accepted_terms_version, accepted_terms_at, allow_data_analysis
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateCommand :one
INSERT INTO commands (
    user_id, command_status
) VALUES (
    $1, $2
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

-- name: ListTransactionsByUserId :many
SELECT * FROM transactions t, commands c
WHERE c.id = t.command_id AND c.user_id = $1
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListIntentsByUserId :many
SELECT * FROM intents i, commands c
WHERE c.id = i.command_id AND c.user_id = $1
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;


-- name: ListCommandsByUserId :many
SELECT * FROM commands c
WHERE c.user_id = $1
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;
