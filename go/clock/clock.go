package clock

import (
	"fmt"
	"time"
)

// Clock defines a type for keeping time.
type Clock struct {
	hour int
	min  int
}

// New is a constructor to create instances of Clock.
func New(hour, min int) Clock {
	t := time.Date(0, 0, 0, hour, min, 0, 0, time.UTC)
	return Clock{hour: t.Hour(), min: t.Minute()}
}

// Add moves the clock forward by the provided number of minutes.
func (clock Clock) Add(minutes int) Clock {
	t := clock.time().Add(durationFromMinutes(minutes))
	return Clock{hour: t.Hour(), min: t.Minute()}
}

// Subtract moves the clock backward by the provided number of minutes.
func (clock Clock) Subtract(minutes int) Clock {
	t := clock.time().Add(-durationFromMinutes(minutes))
	return Clock{hour: t.Hour(), min: t.Minute()}
}

func (clock Clock) String() string {
	return clock.time().Format("15:04")
}

// time converts the clock to a time.Time instance.
func (clock Clock) time() time.Time {
	return time.Date(0, 0, 0, clock.hour, clock.min, 0, 0, time.UTC)
}

func durationFromMinutes(minutes int) time.Duration {
	d, _ := time.ParseDuration(fmt.Sprintf("%dm", minutes)) // WARNING: Ignoring errors
	return d
}
