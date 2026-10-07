package clockadapter

import (
	"library/port/outport"
	"time"
)

type Clock struct{}

func NewClock() outport.Clock {
	return Clock{}
}

func (Clock) Now() time.Time {
	return time.Now()
}
