package chrono

import (
	"testing"
	"time"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/chrono/provider"
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

func TestGetTimeByNanoSeconds(t *testing.T) {
	time.Local = time.UTC
	expected := time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC)

	var ts chrono.TimeStamp = &timeStampMock{}

	got := ts.GetTimeByNanoSeconds(1588291200000000000)

	if expected.String() != got.String() {
		t.Errorf("GetCurrentTime() failed, expected %v, got %v", expected, got)
	}
}

var tsAux chrono.TimeStamp = &provider.TimeStampImpl{}

type timeStampMock struct {
	Time time.Time
}

func (tp *timeStampMock) GetCurrentTime() time.Time {
	return tp.Time
}

func (tp *timeStampMock) GetBlackFridayDay() int {
	return tsAux.GetBlackFridayDay()
}

func (tp *timeStampMock) GetBlackFridayMonth() time.Month {
	return tsAux.GetBlackFridayMonth()
}

func (tp *timeStampMock) GetTimeByNanoSeconds(nanos int64) time.Time {
	return tsAux.GetTimeByNanoSeconds(nanos)
}
