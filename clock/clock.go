package clock

import (
	"time"
)

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (r RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (fc FixedClocker) Now() time.Time {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return time.Date(2023, 2, 10, 12, 34, 56, 0, loc)
}
