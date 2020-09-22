package discountuc

import (
	"time"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/model"
	"github.com/zeroberto/products-store/discount-calculator/model/rule"
)

const (
	// BlackFridayPct represents the discount percentage for blackfriday
	BlackFridayPct float32 = 0.10
	// DefaultPct represents the default discount percentage
	DefaultPct float32 = 0
	// MaxPct represents the maximum discount percentage
	MaxPct float32 = 0.10
	// UserBirthdayPct represents the discount percentage for birthdays
	UserBirthdayPct float32 = 0.05
)

// RuleProviderUseCase is responsible for implementing the DiscountRuleUseCase interface
type RuleProviderUseCase struct {
	TimeStamp chrono.TimeStamp
}

// GetDiscountRules is responsible for providing all discount rules
func (drpuc *RuleProviderUseCase) GetDiscountRules(time time.Time, user *model.User) []model.DiscountRule {
	return []model.DiscountRule{
		&rule.BlackFridayRule{
			Time:             time,
			BlackFridayDay:   drpuc.TimeStamp.GetBlackFridayDay(),
			BlackFridayMonth: drpuc.TimeStamp.GetBlackFridayMonth(),
			BlackFridayPct:   BlackFridayPct,
			DefaultPct:       DefaultPct,
		},
		&rule.UserBirthdayRule{
			Time:            time,
			User:            user,
			UserBirthdayPct: UserBirthdayPct,
			DefaultPct:      DefaultPct,
		},
	}
}
