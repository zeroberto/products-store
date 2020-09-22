package rule

import (
	"time"

	"github.com/zeroberto/products-store/discount-calculator/model"
)

// UserBirthdayRule is responsible for implementing the discountRule interface and
// applying the discount following the birthday rules
type UserBirthdayRule struct {
	Time            time.Time
	User            *model.User
	UserBirthdayPct float32
	DefaultPct      float32
}

// CalculateDiscount is responsible for making a discount calculation
// following the rules of birthday
func (ubr *UserBirthdayRule) CalculateDiscount() float32 {
	if ubr.isBirthday() {
		return ubr.UserBirthdayPct
	}
	return ubr.DefaultPct
}

func (ubr *UserBirthdayRule) isBirthday() bool {
	return ubr.Time.Day() == ubr.User.DateOfBirth.Day() && ubr.Time.Month() == ubr.User.DateOfBirth.Month()
}
