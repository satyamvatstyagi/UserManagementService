package instrumentation

import (
	"context"

	"go.elastic.co/apm/v2"
)

func TraceAPMRequest(ctx context.Context, name string, spanType string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, spanType)
}
