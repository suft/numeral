// A numeral utility package for Go
package numeral

import (
	"fmt"
)

// Version of the numeral library
const Version = "v0.1.0"

// Ordinal Constants
const (
	// The st ordinal suffix
	st = "st"

	// The nd ordinal suffix
	nd = "nd"

	// The rd ordinal suffix
	rd = "rd"

	// The th ordinal suffix
	th = "th"
)

// Ordinal takes a unsigned integer and
// returns the English ordinal letters
// following the inputted integer.
func Ordinal(n uint) string {
	tens := n % 10
	hundreds := n % 100

	// pad the number with the suffix
	pad := func(suff string) string {
		return fmt.Sprintf("%d%s", n, suff)
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
