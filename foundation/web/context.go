package web

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const (
	tracerKey ctxKey = iota + 1
)

func setTraceID(ctx context.Context, traceID uuid.UUID) context.Context {
	return context.WithValue(ctx, tracerKey, traceID)
}

func GetTraceID(ctx context.Context) uuid.UUID {
	v, ok := ctx.Value(tracerKey).(uuid.UUID)
	if !ok {
		return uuid.UUID{}
	}

	return v
}
