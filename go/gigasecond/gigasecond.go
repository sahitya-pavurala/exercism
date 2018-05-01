// Package gigasecond has a function to return the time after adding giga seconds
package gigasecond

import "time"

// AddGigasecond adds and returns the time after giga seconds
func AddGigasecond(t time.Time) time.Time {
	// tt store the time after giga seconds from t
	tt := t.Add(time.Second * 1e9)
	return tt
}
