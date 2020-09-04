package chronofactory

import (
	"github.com/zeroberto/products-store/products-list-endpoint/chrono"
	"github.com/zeroberto/products-store/products-list-endpoint/chrono/provider"
	"github.com/zeroberto/products-store/products-list-endpoint/container"
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
