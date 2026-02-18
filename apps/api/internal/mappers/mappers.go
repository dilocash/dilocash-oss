// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mappers

import (
	"github.com/dilocash/dilocash-oss/apps/api/internal/domain"
	database "github.com/dilocash/dilocash-oss/apps/api/internal/generated/db/postgres"
	transport "github.com/dilocash/dilocash-oss/apps/api/internal/generated/transport/dilocash/v1"
)

// goverter:converter
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:CopyDecimal
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:CopyTime
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:CopyUUID
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:DecimalToString
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:StringToDecimal
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:UUIDToString
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:StringToUUID
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:TimeToTimestamp
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:TimestampToTime
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:PgTextToString
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:StringToPgText
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:PgBoolToBool
// goverter:extend github.com/dilocash/dilocash-oss/apps/api/internal/mappers:BoolToPgBool
type Converter interface {
	// Database -> Domain
	TransactionFromDBToDomain(db database.Transaction) domain.Transaction
	ToDomainUser(db database.User) domain.User

	// Domain -> Database
	ToDBTransaction(d domain.Transaction) database.Transaction
	ToDBCreateTransactionParams(d domain.Transaction) database.CreateTransactionParams
	ToDBUser(d domain.User) database.User

	// Domain -> Transport
	// goverter:ignore state sizeCache unknownFields
	// goverter:map ID Id
	// goverter:map UserID UserId
	ToTransportTransaction(d domain.Transaction) *transport.Transaction

	// Transport -> Domain
	// goverter:ignoreMissing
	// goverter:useZeroValueOnPointerInconsistency
	TransactionFromTransportToDomain(t *transport.CreateTransactionRequest) domain.Transaction
}
