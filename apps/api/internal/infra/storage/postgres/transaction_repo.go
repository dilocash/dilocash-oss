// TODO add license header
package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/your-username/dilocash-oss/apps/api/internal/domain"
)

type transactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) domain.TransactionRepository {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) Save(ctx context.Context, tx *domain.Transaction) error {
	query := `
		INSERT INTO transactions (id, user_id, amount, currency, category, description, raw_input, created_at)
		VALUES (:id, :user_id, :amount, :currency, :category, :description, :raw_input, :created_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, tx)
	return err
}

func (r *transactionRepo) ListByUserID(ctx context.Context, userID string, limit int) ([]*domain.Transaction, error) {
	var txs []*domain.Transaction
	query := `
		SELECT id, user_id, amount, currency, category, description, raw_input, created_at 
		FROM transactions 
		WHERE user_id = $1 
		ORDER BY created_at DESC 
		LIMIT $2
	`
	err := r.db.SelectContext(ctx, &txs, query, userID, limit)
	return txs, err
}
