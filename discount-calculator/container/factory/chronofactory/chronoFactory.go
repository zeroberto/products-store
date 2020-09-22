package chronofactory

import (
	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/chrono/provider"
	"github.com/zeroberto/products-store/discount-calculator/container"
)

// MakeTimeStamp is responsible for providing an instance of TimeStamp
func MakeTimeStamp(c container.Container) chrono.TimeStamp {
	ts, ok := c.Get(container.TimeStampKey)
	if !ok {
		ts = &provider.TimeStampImpl{}
		c.Put(container.TimeStampKey, ts)
	}
	return ts.(chrono.TimeStamp)
}
