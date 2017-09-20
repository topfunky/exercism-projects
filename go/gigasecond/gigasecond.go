package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds one billion seconds to a given time.
// It returns a new time representing the resulting point in time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1e9))
}
