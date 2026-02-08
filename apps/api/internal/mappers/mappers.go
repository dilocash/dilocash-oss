// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mappers

import (
	"github.com/dilocash/dilocash-oss/internal/domain"
	database "github.com/dilocash/dilocash-oss/internal/generated/db/postgres"
	transport "github.com/dilocash/dilocash-oss/internal/generated/transport/dilocash/v1"
)

// goverter:converter
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:CopyDecimal
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:CopyTime
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:CopyUUID
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:DecimalToString
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:StringToDecimal
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:UUIDToString
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:StringToUUID
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:TimeToTimestamp
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:TimestampToTime
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:PgTextToString
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:StringToPgText
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:PgBoolToBool
// goverter:extend github.com/dilocash/dilocash-oss/internal/mappers:BoolToPgBool
type Converter interface {
	// Database -> Domain
	TransactionFromDBToDomain(db database.Transaction) domain.Transaction
	ToDomainUser(db database.User) domain.User

	// Domain -> Database
	ToDBTransaction(d domain.Transaction) database.Transaction
	ToDBUser(d domain.User) database.User

	// Domain -> Transport
	// goverter:map ID TransactionId
	// goverter:ignore state sizeCache unknownFields
	ToTransportTransaction(d domain.Transaction) *transport.ProcessIntentResponse

	// Transport -> Domain
	// goverter:ignoreMissing
	TransactionFromTransportToDomain(t transport.CreateTransactionRequest) *domain.Transaction
}
