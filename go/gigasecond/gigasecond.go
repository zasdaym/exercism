// Solution for problem Gigasecond.
package gigasecond

import "time"

// AddGigasecond will return the given time added with 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
