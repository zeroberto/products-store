package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/model"
	"github.com/zeroberto/products-store/discount-calculator/usecase"
	"github.com/zeroberto/products-store/discount-calculator/usecase/discountuc"
)

func TestRuleProviderUseCaseGetDiscountRules(t *testing.T) {
	expected := 2

	targetTime := time.Date(2020, time.December, 20, 0, 0, 0, 0, time.UTC)
	targetUser := &model.User{DateOfBirth: targetTime}
	var ts chrono.TimeStamp = &timeStampMock{Time: targetTime}

	var druc usecase.DiscountRuleUseCase = &discountuc.RuleProviderUseCase{TimeStamp: ts}

	got := len(druc.GetDiscountRules(targetTime, targetUser))

	if expected != got {
		t.Errorf("GetDiscountRules() failed, expected %v, got %v", expected, got)
	}
}

func TestRuleProviderUseCaseGetDiscountRules_WhenUserIsNil(t *testing.T) {
	expected := 1

	targetTime := time.Date(2020, time.December, 20, 0, 0, 0, 0, time.UTC)
	var targetUser *model.User = nil
	var ts chrono.TimeStamp = &timeStampMock{Time: targetTime}

	var druc usecase.DiscountRuleUseCase = &discountuc.RuleProviderUseCase{TimeStamp: ts}

	got := len(druc.GetDiscountRules(targetTime, targetUser))

	if expected != got {
		t.Errorf("GetDiscountRules() failed, expected %v, got %v", expected, got)
	}
}

func TestRuleCalculatorUseCaseCalculateDiscount_WhenIsBlackFriday_ThenDiscountEqualToTenPercent(t *testing.T) {
	expected := &model.Discount{
		Percentage:   0.10,
		ValueInCents: 11,
	}

	targetProduct := &model.Product{PriceInCents: 110}
	targetUser := &model.User{DateOfBirth: time.Date(2020, time.December, 20, 0, 0, 0, 0, time.UTC)}
	targetTime := time.Date(2020, time.November, 25, 0, 0, 0, 0, time.UTC)
	var ts chrono.TimeStamp = &timeStampMock{Time: targetTime}

	var duc usecase.DiscountUseCase = &discountuc.RuleCalculatorUseCase{
		TimeStamp: ts,
		DRUC:      &discountuc.RuleProviderUseCase{TimeStamp: ts},
	}

	got := duc.CalculateDiscount(targetProduct, targetUser)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

func TestRuleCalculatorUseCaseCalculateDiscount_WhenIsBirthday_ThenDiscountEqualToFivePercent(t *testing.T) {
	expected := &model.Discount{
		Percentage:   0.05,
		ValueInCents: 6,
	}

	targetProduct := &model.Product{PriceInCents: 110}
	targetUser := &model.User{DateOfBirth: time.Date(2020, time.December, 20, 0, 0, 0, 0, time.UTC)}
	targetTime := time.Date(2020, time.December, 20, 0, 0, 0, 0, time.UTC)
	var ts chrono.TimeStamp = &timeStampMock{Time: targetTime}

	var duc usecase.DiscountUseCase = &discountuc.RuleCalculatorUseCase{
		TimeStamp: ts,
		DRUC:      &discountuc.RuleProviderUseCase{TimeStamp: ts},
	}

	got := duc.CalculateDiscount(targetProduct, targetUser)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

func TestRuleCalculatorUseCaseCalculateDiscount_WhenIsBirthdayAndBlackFriday_ThenDiscountEqualToTenPercent(t *testing.T) {
	expected := &model.Discount{
		Percentage:   0.10,
		ValueInCents: 11,
	}

	targetProduct := &model.Product{PriceInCents: 110}
	targetUser := &model.User{DateOfBirth: time.Date(2020, time.November, 25, 0, 0, 0, 0, time.UTC)}
	targetTime := time.Date(2020, time.November, 25, 0, 0, 0, 0, time.UTC)
	var ts chrono.TimeStamp = &timeStampMock{Time: targetTime}

	var duc usecase.DiscountUseCase = &discountuc.RuleCalculatorUseCase{
		TimeStamp: ts,
		DRUC:      &discountuc.RuleProviderUseCase{TimeStamp: ts},
	}

	got := duc.CalculateDiscount(targetProduct, targetUser)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

func TestRuleCalculatorUseCaseCalculateDiscount_WhenIsNotBirthdayNorBlackFriday_ThenDiscountEqualToZeroPercent(t *testing.T) {
	expected := &model.Discount{
		Percentage:   0.00,
		ValueInCents: 0,
	}

	targetProduct := &model.Product{PriceInCents: 110}
	targetUser := &model.User{DateOfBirth: time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC)}
	targetTime := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	var ts chrono.TimeStamp = &timeStampMock{Time: targetTime}

	var duc usecase.DiscountUseCase = &discountuc.RuleCalculatorUseCase{
		TimeStamp: ts,
		DRUC:      &discountuc.RuleProviderUseCase{TimeStamp: ts},
	}

	got := duc.CalculateDiscount(targetProduct, targetUser)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

type timeStampMock struct {
	Time time.Time
}

func (tp *timeStampMock) GetCurrentTime() time.Time {
	return tp.Time
}

func (tp *timeStampMock) GetBlackFridayDay() int {
	return 25
}

func (tp *timeStampMock) GetBlackFridayMonth() time.Month {
	return time.November
}

func (tp *timeStampMock) GetTimeByNanoSeconds(nanos int64) time.Time {
	time.Local = time.UTC
	return time.Unix(0, nanos)
}
