package day1

import (
	"testing"
)

func TestDayOne(t *testing.T) {

	got := Day1("./day1.txt")

	if got != 7 {
		t.Errorf("got != 7, got: %d", got)
	}
}
