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
func Ordinal(n int, suffixOnly bool) (string, error) {
	// Negative number
	if n < 0 {
		return empty, errors.New("negative number found")
	}

	tens := n % 10
	hundreds := n % 100

	// pad the number with the suffix
	pad := func(suff string) string {
		if suffixOnly {
			return suff
		}
		return fmt.Sprintf("%d%s", n, suff)
	}

	switch hundreds {
	case 11, 12, 13:
		return pad(th), nil
	}

	switch tens {
	case 1:
		return pad(st), nil
	case 2:
		return pad(nd), nil
	case 3:
		return pad(rd), nil
	}

	return pad(th), nil
}