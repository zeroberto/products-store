package rule

import (
	"time"
)

// BlackFridayRule is responsible for implementing the discountRule interface and
// applying the discount following the rules of the black friday
type BlackFridayRule struct {
	Time             time.Time
	BlackFridayDay   int
	BlackFridayMonth time.Month
	BlackFridayPct   float32
	DefaultPct       float32
}

// CalculateDiscount is responsible for making a discount calculation
// following the rules of blackfriday
func (bfr *BlackFridayRule) CalculateDiscount() float32 {
	if bfr.isBlackFriday() {
		return bfr.BlackFridayPct
	}
	return bfr.DefaultPct
}

func (bfr *BlackFridayRule) isBlackFriday() bool {
	return bfr.Time.Day() == bfr.BlackFridayDay && bfr.Time.Month() == bfr.BlackFridayMonth
}
