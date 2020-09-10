package chrono

import (
	"time"
)

// TimeStamp is responsible for providing reliable application time
type TimeStamp interface {
	// GetCurrentTime provides date and time of the moment
	GetCurrentTime() time.Time
}
