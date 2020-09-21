package chrono

import (
	"time"
)

const (
	// BlackFridayDay represents the black friday celebration day
	BlackFridayDay int = 25
	// BlackFridayMonth represents the black friday celebration month
	BlackFridayMonth time.Month = time.November
)

// TimeStamp is responsible for providing reliable application time
type TimeStamp interface {
	// GetCurrentTime provides date and time of the moment
	GetCurrentTime() time.Time
	// GetBlackFridayDay provides Black Friday day
	GetBlackFridayDay() int
	// GetBlackFridayMonth provides Black Friday month
	GetBlackFridayMonth() time.Month
}
