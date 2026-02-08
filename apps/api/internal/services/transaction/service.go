package transaction

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	db "github.com/dilocash/dilocash-oss/internal/generated/db/postgres"
	mappers "github.com/dilocash/dilocash-oss/internal/generated/mappers"
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
	req *v1.CreateTransactionRequest,
) (*v1.CreateTransactionResponse, error) {

	// 1. Mapeo de gRPC Request a sqlc Params
	converter := &mappers.ConverterImpl{}
	domainTransaction := converter.TransactionFromTransportToDomain(req)
	dbTransactionParams := converter.ToDBCreateTransactionParams(domainTransaction)

	// 2. Ejecución de la consulta usando sqlc
	transaction, err := s.store.CreateTransaction(ctx, dbTransactionParams)
	if err != nil {
		// Manejo de errores profesional (Logging + gRPC Codes)
		// log.Fatal("failed to store transaction", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to store transaction"))
	}

	domainTransactionResult := converter.TransactionFromDBToDomain(transaction)
	// 4. Mapeo de gRPC Response a Transport Response

	transportTransactionResult := converter.ToTransportTransaction(domainTransactionResult)
	// 3. Respuesta exitosa
	return &v1.CreateTransactionResponse{
		Transaction: transportTransactionResult,
	}, nil
}

func (s *TransactionServer) GetTransaction(
	ctx context.Context,
	req *v1.GetTransactionRequest,
) (*v1.GetTransactionResponse, error) {

	return &v1.GetTransactionResponse{}, errors.New("GetTransaction called. not implemented yet")
}

func (s *TransactionServer) ListTransactions(
	ctx context.Context,
	req *v1.ListTransactionsRequest,
) (*v1.ListTransactionsResponse, error) {

	return &v1.ListTransactionsResponse{}, errors.New("ListTransactions called. not implemented yet")
}
