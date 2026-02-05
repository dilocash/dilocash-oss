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

-- name: CreateTransaction :one
INSERT INTO transactions (
    user_id, amount, currency, category, description, raw_input
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListTransactionsByUserId :many
SELECT * FROM transactions
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
