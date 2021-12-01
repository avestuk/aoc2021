package day1

import "testing"

func TestDay1P2(t *testing.T) {

	got, err := Day1P2("day1_test.txt")
	if err != nil {
		t.Errorf("got err: %s", err)
	}

	if got != 5 {
		t.Errorf("got != 5, got: %d", got)
	}

}

// 199  A
// 200  A B
// 208  A B C
// 210    B C D
// 200  E   C D
// 207  E F   D
// 240  E F G
// 269    F G H
// 260      G H
// 263        H
