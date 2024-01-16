package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger defines the interface for a custom logger
type Logger interface {
	Log(level logrus.Level, message string, fields map[string]interface{})
	Info(message string, fields map[string]interface{})
	Warning(message string, fields map[string]interface{})
	Error(message string, fields map[string]interface{})
	Debug(message string, fields map[string]interface{})
	Fatal(message string, fields map[string]interface{})
	Panic(message string, fields map[string]interface{})
}

// MtnLogger is an implementation of Logger that logs to a file
type MtnLogger struct {
	*logrus.Logger
}

// NewMtnLogger creates a new instance of MtnLogger with log output to a file
func NewMtnLogger(filename string) (*MtnLogger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetOutput(file)

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

	return &MtnLogger{Logger: logger}, nil
}

// Log logs messages with the specified log level and additional fields
func (mtn *MtnLogger) Log(level logrus.Level, message string, fields map[string]interface{}) {
	entry := logrus.NewEntry(mtn.Logger)
	entry.Level = level
	entry.Message = message
	entry.Data = fields
	entry.Logger = mtn.Logger
	entry.Time = time.Now()

	mtn.Logger.WithFields(entry.Data).Log(entry.Level, entry.Message)
}

// Info logs messages with Info level and additional fields
func (mtn *MtnLogger) Info(message string, fields map[string]interface{}) {
	mtn.Log(logrus.InfoLevel, message, fields)
}

// Warning logs messages with Warning level and additional fields
func (mtn *MtnLogger) Warning(message string, fields map[string]interface{}) {
	mtn.Log(logrus.WarnLevel, message, fields)
}

// Error logs messages with Error level and additional fields
func (mtn *MtnLogger) Error(message string, fields map[string]interface{}) {
	mtn.Log(logrus.ErrorLevel, message, fields)
}

// Debug logs messages with Debug level and additional fields
func (mtn *MtnLogger) Debug(message string, fields map[string]interface{}) {
	mtn.Log(logrus.DebugLevel, message, fields)
}

// Fatal logs messages with Fatal level and additional fields
func (mtn *MtnLogger) Fatal(message string, fields map[string]interface{}) {
	mtn.Log(logrus.FatalLevel, message, fields)
}

// Panic logs messages with Panic level and additional fields
func (mtn *MtnLogger) Panic(message string, fields map[string]interface{}) {
	mtn.Log(logrus.PanicLevel, message, fields)
}
