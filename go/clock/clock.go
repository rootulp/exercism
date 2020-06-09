package clock

import (
	"fmt"
	"time"
)

// Clock defines a type for keeping time.
type Clock struct {
	time time.Time
}

// New is a constructor to create instances of Clock.
func New(hour, min int) Clock {
	return Clock{time: time.Date(0, 0, 0, hour, min, 0, 0, time.UTC)}
}

func (c Clock) String() string {
	return c.time.Format("15:04")
}

// Add moves the clock forward by the provided number of minutes.
func (c Clock) Add(minutes int) Clock {
	d, _ := time.ParseDuration(fmt.Sprintf("%dm", minutes)) // WARNING: Ignoring errors
	new := c.time.Add(d)
	return New(new.Hour(), new.Minute())
}

// Subtract moves the clock backward by the provided number of minutes.
func (c Clock) Subtract(minutes int) Clock {
	d, _ := time.ParseDuration(fmt.Sprintf("%dm", minutes)) // WARNING: Ignoring errors
	new := c.time.Add(-d)
	return New(new.Hour(), new.Minute())
}

// It seems strange Go does not include a default implementation of absolute value for ints
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
