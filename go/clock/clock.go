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
	t := asTime(hour, min)
	return Clock{hour: t.Hour(), min: t.Minute()}
}

func (c Clock) String() string {
	return c.time().Format("15:04")
}

// Add moves the clock forward by the provided number of minutes.
func (c Clock) Add(minutes int) Clock {
	duration := durationFromMinutes(minutes)
	new := c.time().Add(duration)
	return New(new.Hour(), new.Minute())
}

// Subtract moves the clock backward by the provided number of minutes.
func (c Clock) Subtract(minutes int) Clock {
	duration := durationFromMinutes(minutes)
	new := c.time().Add(-duration)
	return New(new.Hour(), new.Minute())
}

// time converts the clock to a time.Time instance.
func (c Clock) time() time.Time {
	return asTime(c.hour, c.min)
}

func asTime(hour, min int) time.Time {
	return time.Date(0, 0, 0, hour, min, 0, 0, time.UTC)
}

func durationFromMinutes(minutes int) time.Duration {
	d, _ := time.ParseDuration(fmt.Sprintf("%dm", minutes)) // WARNING: Ignoring errors
	return d
}
