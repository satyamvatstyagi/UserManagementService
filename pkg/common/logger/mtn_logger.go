package logger

import (
	"context"
	"sync"
	"time"

	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger defines the interface for a custom logger
type Logger interface {
	Log(level logrus.Level, message string)
	Info(message string)
	Warning(message string)
	Error(message string)
	Debug(message string)
	Fatal(message string)
	Panic(message string)
	WithContext(ctx context.Context) Logger
	WithFields(fields map[string]interface{}) Logger
	SetLevel(level string)
}

// MtnLogger is an implementation of Logger that logs to a file
type MtnLogger struct {
	*logrus.Logger
	logObject *logObject
}

type logObject struct {
	ctxt   context.Context
	fields map[string]interface{}
	mu     *sync.RWMutex
}

// NewMtnLogger creates a new instance of MtnLogger with log output to a file
func NewMtnLogger(filename string) (Logger, error) {

	fileLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    consts.LoggerFileMaxSize,   // Max size in megabytes before log rotation
		MaxBackups: consts.LoggerFileMaxBackup, // Max number of old log files to retain
		MaxAge:     consts.LoggerFileMaxAge,    // Max number of days to retain old log files
		LocalTime:  true,                       // Use local time for log rotation
		Compress:   false,                      // Compress old log files
	}
	logger := logrus.New()
	logger.SetOutput(fileLogger)

	// Use JSON formatter
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "ts",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
			logrus.FieldKeyFunc:  "caller",
		},
	})

	logObj := &logObject{
		fields: make(map[string]interface{}),
		mu:     &sync.RWMutex{},
	}

	return &MtnLogger{Logger: logger, logObject: logObj}, nil
}

func (mtn *MtnLogger) WithContext(ctx context.Context) Logger {
	mtn.logObject.ctxt = ctx
	return mtn
}

// WithFields implementation
func (mtn *MtnLogger) WithFields(fields map[string]interface{}) Logger {
	mtn.logObject.fields = fields
	return mtn
}

func (mtn *MtnLogger) Context() context.Context {
	return mtn.logObject.ctxt
}

func (mtn *MtnLogger) Fields() map[string]interface{} {
	return mtn.logObject.fields
}

// Log logs messages with the specified log level and additional fields
func (mtn *MtnLogger) Log(level logrus.Level, message string) {

	mtn.logObject.mu.Lock()
	defer mtn.logObject.mu.Unlock()

	fields := mtn.Fields()
	ctx := mtn.Context()

	if ctx != nil {
		if ctxFields, ok := ctx.Value(consts.LogContext).(map[string]interface{}); ok {
			for key, value := range ctxFields {
				fields[key] = value
			}
		}
	}
	// Add the service name to the fields
	fields["service_name"] = consts.AppName

	// Create a new log entry
	entry := logrus.NewEntry(mtn.Logger)
	entry.Data = fields
	entry.Message = message
	entry.Level = level
	entry.Logger = mtn.Logger
	entry.Time = time.Now()

	mtn.Logger.WithFields(entry.Data).Log(entry.Level, entry.Message)
	mtn.WithFields(make(map[string]interface{}))
}

// Info logs messages with Info level and additional fields
func (mtn *MtnLogger) Info(message string) {
	mtn.Log(logrus.InfoLevel, message)
}

// Warning logs messages with Warning level and additional fields
func (mtn *MtnLogger) Warning(message string) {
	mtn.Log(logrus.WarnLevel, message)
}

// Error logs messages with Error level and additional fields
func (mtn *MtnLogger) Error(message string) {
	mtn.Log(logrus.ErrorLevel, message)
}

// Debug logs messages with Debug level and additional fields
func (mtn *MtnLogger) Debug(message string) {
	mtn.Log(logrus.DebugLevel, message)
}

// Fatal logs messages with Fatal level and additional fields
func (mtn *MtnLogger) Fatal(message string) {
	mtn.Log(logrus.FatalLevel, message)
}

// Panic logs messages with Panic level and additional fields
func (mtn *MtnLogger) Panic(message string) {
	mtn.Log(logrus.PanicLevel, message)
}

// SetLevel sets the log level for the logger
func (mtn *MtnLogger) SetLevel(level string) {
	defaultLevel := logrus.InfoLevel
	switch level {
	case logrus.TraceLevel.String():
		defaultLevel = logrus.TraceLevel
	case logrus.DebugLevel.String():
		defaultLevel = logrus.DebugLevel
	case logrus.InfoLevel.String():
		defaultLevel = logrus.InfoLevel
	case logrus.WarnLevel.String():
		defaultLevel = logrus.WarnLevel
	case logrus.ErrorLevel.String():
		defaultLevel = logrus.ErrorLevel
	case logrus.FatalLevel.String():
		defaultLevel = logrus.FatalLevel
	case logrus.PanicLevel.String():
		defaultLevel = logrus.PanicLevel
	}
	mtn.Logger.SetLevel(defaultLevel)
}
