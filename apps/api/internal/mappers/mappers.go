// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mappers

import (
	domain "github.com/dilocash/dilocash-oss/apps/api/internal/domain"
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
	CommandRowFromDBToDB(db database.GetCommandsSyncRow) database.Command
	CommandFromDBToDomain(db database.Command) *domain.Command
	IntentRowFromDBToDB(db database.GetIntentsSyncRow) database.Intent
	// goverter:useZeroValueOnPointerInconsistency
	IntentFromDBToDomain(db database.Intent) *domain.Intent
	TransactionRowFromDBToDB(db database.GetTransactionsSyncRow) database.Transaction
	// goverter:useZeroValueOnPointerInconsistency
	TransactionFromDBToDomain(db database.Transaction) *domain.Transaction
	// goverter:useZeroValueOnPointerInconsistency
	ProfileFromDBToDomain(db database.Profile) *domain.Profile

	// Domain -> Database
	// goverter:useZeroValueOnPointerInconsistency
	ToDBTransaction(d *domain.Transaction) database.Transaction
	// goverter:useZeroValueOnPointerInconsistency
	ToDBCommand(d *domain.Command) database.Command
	// goverter:useZeroValueOnPointerInconsistency
	ToDBIntent(d *domain.Intent) database.Intent

	// Domain -> Database params
	// goverter:useZeroValueOnPointerInconsistency
	ToDBCreateCommandParams(d *domain.Command) database.CreateCommandParams
	// goverter:useZeroValueOnPointerInconsistency
	ToDBUpdateCommandParams(d *domain.Command) database.UpdateCommandParams
	// goverter:useZeroValueOnPointerInconsistency
	ToDBCreateIntentParams(d *domain.Intent) database.CreateIntentParams
	// goverter:useZeroValueOnPointerInconsistency
	ToDBUpdateIntentParams(d *domain.Intent) database.UpdateIntentParams
	// goverter:useZeroValueOnPointerInconsistency
	ToDBCreateTransactionParams(d *domain.Transaction) database.CreateTransactionParams
	// goverter:useZeroValueOnPointerInconsistency
	ToDBUpdateTransactionParams(d *domain.Transaction) database.UpdateTransactionParams

	// Domain -> Transport
	// goverter:ignore state sizeCache unknownFields
	// goverter:map ID Id
	// goverter:map CommandID CommandId
	ToTransportTransaction(d *domain.Transaction) *transport.Transaction
	// goverter:map ID Id
	// goverter:ignore state sizeCache unknownFields
	// goverter:enum no
	ToTransportCommand(d *domain.Command) *transport.Command
	// goverter:map ID Id
	// goverter:map CommandID CommandId
	// goverter:ignore state sizeCache unknownFields
	ToTransportIntent(d *domain.Intent) *transport.Intent

	// Transport -> Domain
	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Id ID
	// goverter:ignore ProfileID
	CommandFromTransportToDomain(t *transport.Command) *domain.Command
	// goverter:ignoreMissing
	// goverter:useZeroValueOnPointerInconsistency
	IntentFromTransportToDomain(t *transport.Intent) *domain.Intent
	// goverter:ignoreMissing
	// goverter:useZeroValueOnPointerInconsistency
	TransactionFromTransportToDomain(t *transport.Transaction) *domain.Transaction
}
