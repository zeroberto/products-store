package provider

import (
	"time"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
)

// TimeStampImpl is Responsible for implementing the methods that provide
// the application time reliably
type TimeStampImpl struct{}

// GetCurrentTime provides date and time of the moment
func (tp *TimeStampImpl) GetCurrentTime() time.Time {
	return time.Now()
}

// GetBlackFridayDay provides Black Friday day
func (tp *TimeStampImpl) GetBlackFridayDay() int {
	return chrono.BlackFridayDay
}

// GetBlackFridayMonth provides Black Friday month
func (tp *TimeStampImpl) GetBlackFridayMonth() time.Month {
	return chrono.BlackFridayMonth
}

// GetTimeByNanoSeconds obtains a time from the nanoseconds
func (tp *TimeStampImpl) GetTimeByNanoSeconds(nanos int64) time.Time {
	time.Local = time.UTC
	return time.Unix(0, nanos)
}
