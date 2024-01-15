package consts

import "time"

const KeySeperator = "."

const (
	DefaultExpiration = 5 * time.Minute
	PurgeTime         = 10 * time.Minute
	MaxTimeout        = 25 * time.Second
)

const (
	UniqueViolation     = "23505"
	ForeignKeyViolation = "23503"
)
