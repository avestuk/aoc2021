package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1P2(t *testing.T) {

	got, err := Day1P22("day1_test.txt")
	if err != nil {
		t.Errorf("got err: %s", err)
	}

	if got != 5 {
		t.Errorf("got != 5, got: %d", got)
	}

}

func TestPopulate(t *testing.T) {
	w := &window{}

	w.populate(0)
	require.Equal(t, 0, w.measurements[0])
	w.populate(1)
	require.Equal(t, 1, w.measurements[1])
	w.populate(2)
	require.Equal(t, 2, w.measurements[2])

	w.populate(3)
	require.Equal(t, 3, w.measurements[2])
	require.Equal(t, 2, w.measurements[1])
	require.Equal(t, 1, w.measurements[0])
}

func TestIncrease(t *testing.T) {
	w := &window{}

	// First Frame
	w.populate(0)
	w.populate(1)
	w.populate(2)

	// Second Frame
	w.populate(3)

	require.Equal(t, 1, w.increase)

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
