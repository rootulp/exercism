package gigasecond

import (
	"time"
)

const Gigasecond = 1000000000 * time.Second

// AddGigasecond returns the time provided plus one gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
