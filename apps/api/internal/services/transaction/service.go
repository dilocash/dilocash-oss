package transaction

import (
	"context"
	"errors"

	db "github.com/dilocash/dilocash-oss/internal/generated/db/postgres"
	v1 "github.com/dilocash/dilocash-oss/internal/generated/transport/dilocash/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TransactionServer implementa la interfaz generada por gRPC/Connect
type TransactionServer struct {
	store *db.Queries
	pool  *pgxpool.Pool
}

func NewTransactionServer(pool *pgxpool.Pool) *TransactionServer {
	return &TransactionServer{
		store: db.New(pool), // sqlc.New acepta un DB TX o un Pool
		pool:  pool,
	}
}

func (s *TransactionServer) CreateTransaction(
	ctx context.Context,
	req *connect.Request[v1.CreateTransactionRequest],
) (*connect.Response[v1.CreateTransactionResponse], error) {

	// 1. Mapeo de gRPC Request a sqlc Params
	arg := db.CreateTransactionParams{
		UserID:      req.Msg.UserId,
		Amount:      req.Msg.Amount,
		Currency:    req.Msg.Currency,
		Category:    req.Msg.Category,
		Description: req.Msg.Description,
		RawInput:    req.Msg.RawInput,
	}

	// 2. Ejecución de la consulta usando sqlc
	transaction, err := s.store.CreateTransaction(ctx, arg)
	if err != nil {
		// Manejo de errores profesional (Logging + gRPC Codes)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to store transaction"))
	}

	// 3. Respuesta exitosa
	return connect.NewResponse(&v1.CreateTransactionResponse{
		Id:        transaction.ID.String(),
		CreatedAt: transaction.CreatedAt.String(),
	}), nil
}
