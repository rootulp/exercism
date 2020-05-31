package gigasecond

import (
	"time"
)

const Gigasecond = 1e9 * time.Second

// AddGigasecond returns the time provided plus one gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
