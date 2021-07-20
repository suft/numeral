package numeral

import (
	"errors"
	"fmt"
)

// Version of the numeral library
const Version = "v0.1.0"

// Ordinal Constants
const (
	// The empty ordinal suffix
	empty = ""

	// The st ordinal suffix
	st = "st"

	// The nd ordinal suffix
	nd = "nd"

	// The rd ordinal suffix
	rd = "rd"

	// The th ordinal suffix
	th = "th"
)

// Ordinal returns the ordinal value of an integer
func Ordinal(n int) (string, error) {
	// Negative number
	if n < 0 {
		return empty, errors.New("negative number found")
	}

	tens := n % 10
	hundreds := n % 100

	// pad the number with the suffix
	pad := func(suff string) (string, error) {
		return fmt.Sprintf("%d%s", n, suff), nil
	}

	switch hundreds {
	case 11, 12, 13:
		return pad(th)
	}

	switch tens {
	case 1:
		return pad(st)
	case 2:
		return pad(nd)
	case 3:
		return pad(rd)
	}

	return pad(th)
}
