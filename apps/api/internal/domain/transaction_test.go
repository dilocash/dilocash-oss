// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTransactionDomainStructs(t *testing.T) {

	tests := []struct {
		name        string
		transaction Transaction
		wantErr     bool // true if we expect a validation error
	}{
		{
			name: "Valid Transaction",
			transaction: Transaction{
				ID:        uuid.New(),
				Amount:    decimal.NewFromInt(1000),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				CommandID: uuid.New(),
				Currency:  "USD",
			},
			wantErr: false,
		},
		{
			name: "Invalid sync information",
			transaction: Transaction{
				ID:       uuid.New(),
				Amount:   decimal.NewFromInt(1000),
				Currency: "USD",
			},
			wantErr: true,
		},
		{
			name: "No amount",
			transaction: Transaction{
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "Invalid zero Amount",
			transaction: Transaction{
				ID:     uuid.New(),
				Amount: decimal.NewFromInt(0),
			},
			wantErr: true,
		},
		{
			name: "Invalid  negative Amount",
			transaction: Transaction{
				ID:     uuid.New(),
				Amount: decimal.NewFromInt(-1),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.transaction) // The core validation call
			if (err != nil) != tt.wantErr {
				t.Errorf("validate.Struct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	t.Run("Transaction", func(t *testing.T) {
		id := uuid.New()
		tx := &Transaction{
			ID:       id,
			Currency: "USD",
		}
		assert.Equal(t, id, tx.ID)
		assert.Equal(t, "USD", tx.Currency)
	})
}
