package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          uuid.UUID
	Amount      decimal.Decimal
	Currency    string
	Category    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
	CommandID   uuid.UUID
}

type TransactionsSync struct {
	Created []Transaction
	Updated []Transaction
	Deleted []uuid.UUID
}

type TransactionRepository interface {
	PullTransactionChanges(profileId string, lastPulledAt time.Time) (*TransactionsSync, error)
	PushTransactionChanges(profileId string, transactionsSync *TransactionsSync) error
}
