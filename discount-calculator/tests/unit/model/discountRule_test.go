package model

import (
	"testing"
	"time"

	"github.com/zeroberto/products-store/discount-calculator/model"
	"github.com/zeroberto/products-store/discount-calculator/model/rule"
)

func TestBlackFridayDiscountRuleCalculateDiscount_WhenIsBlackFriday(t *testing.T) {
	var expected float32 = 0.10

	var rule model.DiscountRule = &rule.BlackFridayRule{
		Time:             time.Date(2020, time.November, 25, 0, 0, 0, 0, time.UTC),
		BlackFridayDay:   25,
		BlackFridayMonth: time.November,
		BlackFridayPct:   0.10,
		DefaultPct:       0.00,
	}

	got := rule.CalculateDiscount()

	if expected != got {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestBlackFridayDiscountRuleCalculateDiscount_WhenIsNotBlackFriday(t *testing.T) {
	var expected float32 = 0.00

	var rule model.DiscountRule = &rule.BlackFridayRule{
		Time:             time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
		BlackFridayDay:   25,
		BlackFridayMonth: time.November,
		BlackFridayPct:   0.10,
		DefaultPct:       0.00,
	}

	got := rule.CalculateDiscount()

	if expected != got {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestUserBirthdayDiscountRuleCalculateDiscount_WhenIsBirthday(t *testing.T) {
	var expected float32 = 0.05

	var rule model.DiscountRule = &rule.UserBirthdayRule{
		Time:            time.Date(2020, time.May, 30, 0, 0, 0, 0, time.UTC),
		User:            &model.User{DateOfBirth: time.Date(2020, time.May, 30, 0, 0, 0, 0, time.UTC)},
		UserBirthdayPct: 0.05,
		DefaultPct:      0.00,
	}

	got := rule.CalculateDiscount()

	if expected != got {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestUserBirthdayDiscountRuleCalculateDiscount_WhenIsNotBirthday(t *testing.T) {
	var expected float32 = 0.00

	var rule model.DiscountRule = &rule.UserBirthdayRule{
		Time:            time.Date(2020, time.May, 30, 0, 0, 0, 0, time.UTC),
		User:            &model.User{DateOfBirth: time.Date(2020, time.May, 31, 0, 0, 0, 0, time.UTC)},
		UserBirthdayPct: 0.05,
		DefaultPct:      0.00,
	}

	got := rule.CalculateDiscount()

	if expected != got {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}
