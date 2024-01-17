package utils

import (
	"context"

	"go.elastic.co/apm/v2"
)

func NewSpan(ctx context.Context, name string, spanType string) (*apm.Span, context.Context) {

	span, ctxt := apm.StartSpan(ctx, name, spanType)
	defer span.End()

	return span, ctxt
}
