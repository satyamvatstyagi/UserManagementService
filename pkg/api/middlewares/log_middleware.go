package middlewares

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
)

// LoggingMiddleware is a middleware that logs the request and response
func LoggingMiddleware(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		trace := true
		fields := map[string]interface{}{
			consts.TraceID: c.Request.Header.Get(consts.TraceID),
			"method":       c.Request.Method,
			"endpoint":     c.Request.URL.Path,
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, consts.LogContext, fields)
		c.Request = c.Request.WithContext(ctx)

		defer func() {
			if trace {
				traceMsg := fmt.Sprintf("Stacktrace: %v", string(debug.Stack()))
				logger.
					WithContext(ctx).
					Panic(traceMsg)
			}
		}()
		// Log the inbound request
		logger.WithContext(ctx).Info("Inbound request")
		timeStart := time.Now()
		c.Next()

		// Check the status code of the response and log accordingly
		statusCode := c.Writer.Status()
		fields["status_code"] = statusCode
		fields["duration_ms"] = time.Since(timeStart).String()
		if statusCode >= 400 {
			logger.WithFields(fields).Error("Outbound response")
		} else {
			logger.WithFields(fields).Info("Outbound response")
		}
		trace = false
	}
}
