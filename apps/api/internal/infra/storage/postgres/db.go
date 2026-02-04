// TODO add license header
package postgres

import (
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Standard library compatibility
	"github.com/jmoiron/sqlx"
)

// NewConnection creates a new SQLX database connection pool
func NewConnection(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		return nil, err
	}

	// Staff-level tuning: Connection pooling
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
