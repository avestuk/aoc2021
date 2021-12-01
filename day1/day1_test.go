package day1

import (
	"testing"
)

func TestDayOne(t *testing.T) {

	got, err := Day1P1("./day1_test.txt")
	if err != nil {
		t.Errorf("got err: %s", err)
	}

	if got != 7 {
		t.Errorf("got != 7, got: %d", got)
	}
}
