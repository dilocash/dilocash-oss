// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// internal/infra/repository/transactor.go
package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTransactor struct {
	pool *pgxpool.Pool
}

func NewPostgresTransactor(pool *pgxpool.Pool) *PostgresTransactor {
	return &PostgresTransactor{
		pool: pool,
	}
}

func (t *PostgresTransactor) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	// Ejecutamos la lógica del Usecase pasando el TX en el contexto
	if err := fn(injectTx(ctx, tx)); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
