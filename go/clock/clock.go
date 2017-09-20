// Package clock provides clock manipulation functions.
package clock

import (
	"fmt"
)

const testVersion = 4

// Clock is a struct for tracking time in hours and minutes.
type Clock struct {
	hour, minute int
}

// New creates a Clock from an hour and minute.
// It returns a Clock struct.
func New(hour, minute int) Clock {
	h, m := rollover(hour, minute)
	return Clock{h, m}
}

// String formats a clock to a readable representation.
// It returns a string representation of a Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add does math on a Clock, adding minutes.
// It returns a Clock struct.
func (c Clock) Add(minutes int) Clock {
	h, m := rollover(c.hour, c.minute+minutes)
	return Clock{h, m}
}

const (
	minutesInHour = 60
	minutesInDay  = (hoursInDay * minutesInHour)
	hoursInDay    = 24
)

// rollover keeps hours within 0-23 and minutes within 0-59.
func rollover(hour, minute int) (h, m int) {
	minutesOnly := hour*minutesInHour + minute
	if minutesOnly < 0 {
		minutesOnly = minutesInDay + (minutesOnly % minutesInDay)
	}

	h = minutesOnly / minutesInHour
	m = minutesOnly % minutesInHour
	if h == hoursInDay {
		h = 0
	} else if h > hoursInDay {
		h = h % hoursInDay
	}
	return
}
