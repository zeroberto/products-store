package provider

import (
	"time"
)

// TimeStampImpl is Responsible for implementing the methods that provide
// the application time reliably
type TimeStampImpl struct{}

// GetCurrentTime provides date and time of the moment
func (tp *TimeStampImpl) GetCurrentTime() time.Time {
	return time.Now()
}
