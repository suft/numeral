package numeral_test

import (
	"testing"

	"github.com/suft/numeral"
)

func TestOrdinalNegative(t *testing.T) {
	_, got := numeral.Ordinal(-1)
	want := "negative number found"

	if got.Error() != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalZero(t *testing.T) {
	got, _ := numeral.Ordinal(0)
	want := "0th"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalFirst(t *testing.T) {
	got, _ := numeral.Ordinal(1)
	want := "1st"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalSecond(t *testing.T) {
	got, _ := numeral.Ordinal(2)
	want := "2nd"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalThird(t *testing.T) {
	got, _ := numeral.Ordinal(3)
	want := "3rd"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalEleventh(t *testing.T) {
	got, _ := numeral.Ordinal(11)
	want := "11th"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestOrdinalTwentyFirst(t *testing.T) {
	got, _ := numeral.Ordinal(21)
	want := "21st"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}
