// TODO add license header
package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          string
	UserID      string
	Amount      decimal.Decimal
	Currency    string // ISO 4217
	Category    string
	Description string
	RawInput    string // Original voice transcript/text
	CreatedAt   time.Time
}

// TransactionRepository defines the storage contract for the ledger.
type TransactionRepository interface {
	Save(ctx context.Context, tx *Transaction) error
	ListByUserID(ctx context.Context, userID string, limit int) ([]*Transaction, error)
}
