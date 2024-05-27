package epoch

import "time"

type Epoch struct {
	time.Time
}

// Creates a new Epoch with the given UNIX millisecond timestamp.
func New(epochUnixMillis int64) *Epoch {
	return &Epoch{Time: time.UnixMilli(epochUnixMillis)}
}

func (e *Epoch) GenerateEpochTimeMillis() int64 {
	return time.Since(e.Time).Milliseconds()
}
