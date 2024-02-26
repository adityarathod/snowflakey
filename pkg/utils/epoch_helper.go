package utils

import "time"

type SnowflakeEpoch struct {
	// Absolute time of this snowflake epoch
	epochTime time.Time
}

// NewSnowflakeEpoch creates a new SnowflakeEpoch with the given UNIX millisecond timestamp.
func NewSnowflakeEpoch(epochUnixMillis int64) *SnowflakeEpoch {
	return &SnowflakeEpoch{epochTime: time.UnixMilli(epochUnixMillis)}
}

func (e *SnowflakeEpoch) GenerateEpochTimeMillis() int64 {
	return time.Since(e.epochTime).Milliseconds()
}
