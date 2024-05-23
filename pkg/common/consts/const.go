package consts

import "time"

const KeySeperator = "."

type contextKey string

const (
	AppName                      = "UserManagementService"
	DefaultExpiration            = 5 * time.Minute
	PurgeTime                    = 10 * time.Minute
	MaxTimeout                   = 25 * time.Second
	AppVersion                   = "1.0.0"
	LogContext        contextKey = "log:context"
	ConfigContext     contextKey = "config:context"
	TraceID                      = "traceID"
)

const (
	UniqueViolation     = "23505"
	ForeignKeyViolation = "23503"
)

const (
	SpanTypeCustum         = "custum"
	SpanTypeQueryExecution = "psql"
)

const (
	LoggerFileMaxSize   = 10  // Max size in megabytes
	LoggerFileMaxAge    = 15  // Max age in days
	LoggerFileMaxBackup = 100 // Maximum number of old log files to retain.
	LoggerDirectory     = "log"
)
