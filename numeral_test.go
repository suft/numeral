package numeral_test

import (
	"testing"

	"github.com/suft/numeral"
)

func TestOrdinalNegativeSuffixOnly(t *testing.T) {
	_, got := numeral.Ordinal(-1, true)
	want := "negative number found"

	if got.Error() != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalNegative(t *testing.T) {
	_, got := numeral.Ordinal(-1, false)
	want := "negative number found"

	if got.Error() != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalZeroSuffixOnly(t *testing.T) {
	got, _ := numeral.Ordinal(0, true)
	want := "th"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalZero(t *testing.T) {
	got, _ := numeral.Ordinal(0, false)
	want := "0th"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalFirstSuffixOnly(t *testing.T) {
	got, _ := numeral.Ordinal(1, true)
	want := "st"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalFirst(t *testing.T) {
	got, _ := numeral.Ordinal(1, false)
	want := "1st"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalSecondSuffixOnly(t *testing.T) {
	got, _ := numeral.Ordinal(2, true)
	want := "nd"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalSecond(t *testing.T) {
	got, _ := numeral.Ordinal(2, false)
	want := "2nd"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalThirdSuffixOnly(t *testing.T) {
	got, _ := numeral.Ordinal(3, true)
	want := "rd"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalThird(t *testing.T) {
	got, _ := numeral.Ordinal(3, false)
	want := "3rd"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalEleventhSuffixOnly(t *testing.T) {
	got, _ := numeral.Ordinal(11, true)
	want := "th"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalEleventh(t *testing.T) {
	got, _ := numeral.Ordinal(11, false)
	want := "11th"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalTwentyFirstSuffixOnly(t *testing.T) {
	got, _ := numeral.Ordinal(21, true)
	want := "st"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalTwentyFirst(t *testing.T) {
	got, _ := numeral.Ordinal(21, false)
	want := "21st"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}
