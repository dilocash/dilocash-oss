// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mappers

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CopyDecimal(d decimal.Decimal) decimal.Decimal {
	return d
}

func CopyTime(t time.Time) time.Time {
	return t
}

func CopyUUID(u uuid.UUID) uuid.UUID {
	return u
}

func DecimalToString(d decimal.Decimal) string {
	return d.String()
}

func StringToDecimal(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}

func UUIDToString(u uuid.UUID) string {
	return u.String()
}

func StringToUUID(s string) uuid.UUID {
	u, _ := uuid.Parse(s)
	return u
}

func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func TimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

func PgTextToString(t pgtype.Text) string {
	if !t.Valid {
		return ""
	}
	return t.String
}

func StringToPgText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{String: *s, Valid: *s != ""}
}

func PgBoolToBool(b pgtype.Bool) bool {
	return b.Bool
}

func BoolToPgBool(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}

func Ptr[T any](v T) *T {
	return &v
}
