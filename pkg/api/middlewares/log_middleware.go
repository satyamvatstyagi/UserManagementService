package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	l "github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
)

// LoggingMiddleware is a middleware that logs the request and response
func LoggingMiddleware(logger l.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		swriter := &statusWriter{c.Writer, 0, 0}
		c.Next()

		fields := map[string]interface{}{
			"endpoint": c.Request.URL.Path,
			"method":   c.Request.Method,
		}

		if swriter.status >= 400 {
			fields["error"] = fmt.Sprintf("Inbound request failed with status %d", swriter.status)
			logger.Error(fields)
		} else {
			fields["message"] = "Inbound request successful"
			logger.Info(fields)
		}
	}
}

// statusWriter is a custom ResponseWriter that captures the HTTP status code.
type statusWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}
