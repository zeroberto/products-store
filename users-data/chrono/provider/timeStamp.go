package provider

import (
	"os"
	"time"
)

// TimeStampImpl is Responsible for implementing the methods that provide
// the application time reliably
type TimeStampImpl struct{}

// GetCurrentTime provides date and time of the moment
func (tp *TimeStampImpl) GetCurrentTime() time.Time {
	if defaultLocaltime := os.Getenv("DEFAULT_LOCALTIME"); defaultLocaltime != "" {
		if localtime, err := time.Parse(time.RFC3339, defaultLocaltime); err == nil {
			return localtime
		}
	}
	return time.Now()
}
