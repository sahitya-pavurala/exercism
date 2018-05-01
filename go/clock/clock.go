// Package clock has methods and functions to handle time without dates
package clock

import (
	"fmt"
)

// Clock is a struct with h hours and m minutes
type Clock struct {
	h int
	m int
}

// New is called to create a new Clock
func New(h int, m int) Clock {
	m, addHours := normalizeMinutes(m)
	h = normalizeHours(h + addHours)
	return Clock{h, m}
}

// String method is used to give the string representation of Clock
func (c Clock) String() string {

	m, addHours := normalizeMinutes(c.m)
	h := normalizeHours(c.h + addHours)

	hrs := fmt.Sprintf("%v", h)
	mins := fmt.Sprintf("%v", m)

	if len(hrs) != 2 {
		hrs = "0" + hrs
	}

	if len(mins) != 2 {
		mins = "0" + mins
	}

	return hrs + ":" + mins

}

// Add method adds the specified number of minutes to the clock
func (c Clock) Add(addMin int) Clock {
	min, addHours := normalizeMinutes(addMin)
	return Clock{c.h + addHours, c.m + min}
}

//Subtract method subtracts the specified number of minutes from the clock
func (c Clock) Subtract(minusMin int) Clock {
	return c.Add(-1 * minusMin)
}

// normalizeMinutes function gives minutes after calculating overflow
func normalizeMinutes(min int) (int, int) {
	flag := false
	addHours := 0
	if min < 0 {
		flag = !flag
		min = -1 * min
	}

	if min >= 60 {
		addHours = min / 60
		min = min % 60
	}

	if flag {
		if min > 0 {
			min = 60 - min
			addHours += 1
		}
		addHours = -1 * addHours
	}

	return min, addHours
}

//// normalizeHours function gives minutes after calculating overflow
func normalizeHours(hrs int) int {
	flag := false

	if hrs < 0 {
		flag = !flag
		hrs = -1 * hrs
	}

	hrs = hrs % 24

	if flag && hrs > 0 {
		hrs = 24 - hrs
	}

	return hrs
}
