package chrono

import (
	"testing"
	"time"

	"github.com/zeroberto/products-store/users-data/chrono"
)

func TestGetCurrentTime(t *testing.T) {
	expected := time.Now()

	var ts chrono.TimeStamp = &timeStampMock{Time: expected}

	got := ts.GetCurrentTime()

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
