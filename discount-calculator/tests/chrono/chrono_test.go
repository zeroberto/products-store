package chrono

import (
	"testing"
	"time"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
)

func TestGetCurrentTime(t *testing.T) {
	expected := time.Now()

	var ts chrono.TimeStamp = &timeStampMock{Time: expected}

	got := ts.GetCurrentTime()

	if expected != got {
		t.Errorf("GetCurrentTime() failed, expected %v, got %v", expected, got)
	}
}

func TestGetBlackFridayDay(t *testing.T) {
	expected := 25

	var ts chrono.TimeStamp = &timeStampMock{}

	got := ts.GetBlackFridayDay()

	if expected != got {
		t.Errorf("GetCurrentTime() failed, expected %v, got %v", expected, got)
	}
}

func TestGetBlackFridayMonth(t *testing.T) {
	expected := time.November

	var ts chrono.TimeStamp = &timeStampMock{}

	got := ts.GetBlackFridayMonth()

	if expected != got {
		t.Errorf("GetCurrentTime() failed, expected %v, got %v", expected, got)
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
