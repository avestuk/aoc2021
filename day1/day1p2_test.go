package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddMeasurement(t *testing.T) {
	// Check that we can add measurements
	measurements := []int{}
	measurements = append(measurements, 0, 1, 2)
	require.Equal(t, 0, measurements[0])
	require.Equal(t, 1, measurements[2])
	require.Equal(t, 2, measurements[2])

	// What happens when we add a fourth measurement?
	measurements = append(measurements, 3)
	// We would expect the window to "slide" such that the previous first
	// element is now the zeroth element
	require.Equal(t, 1, measurements[0])
	require.Equal(t, 3, measurements[2])
}

// When we add a fourth measurement does the window "slide"?
func TestPopulate(t *testing.T) {
	w := &window{}

	w.Populate(0)
	require.Equal(t, 0, w.measurements[0])
	w.Populate(1)
	require.Equal(t, 1, w.measurements[1])
	w.Populate(2)
	require.Equal(t, 2, w.measurements[2])

	// What happens when we add a fourth measurement?
	w.Populate(3)
	require.Equal(t, 3, w.measurements[2])
	// Did the window slide?
	require.Equal(t, 1, w.measurements[0])
}

// When the window slides we need to know whether the new window has a greater
// total than the previous window
func TestSum(t *testing.T) {
	w := &window{}

	w.Populate(0)
	w.Populate(1)
	w.Populate(2)

	got := w.Sum()
	require.Equal(t, 3, got)

	// What happens when we add another element and sum?
	w.Populate(3)
	got = w.Sum()
	require.Equal(t, 6, got)
}

// Now that we have the window sliding and we can calculate window sums we need
// to compare the current window total against the previous total
func TestCompare(t *testing.T) {
	w := &window{}

	// Validate that the previousWindow is set when we hit 3 items
	w.Populate(0)
	w.Populate(1)
	w.Populate(2)
	w.Compare()
	require.Equal(t, 3, w.previousWindow)

	// Validate that the previousWindow is updated
	// and that the
	w.Populate(3)
	w.Compare()
	require.Equal(t, 6, w.previousWindow)
	require.Equal(t, 1, w.increase)
}

// Now we can wire everything together
func TestMain(t *testing.T) {
	got := main("./day1_test.txt")

	require.Equal(t, 5, got)
}

// Answer for the below is 5
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
