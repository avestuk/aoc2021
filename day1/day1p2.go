package day1

import (
	"fmt"
	"strconv"
)

type window struct {
	measurements   []int
	increase       int
	previousWindow int
}

func (w *window) Populate(i int) {
	// If we have less than 3 measurements in the window then append
	if len(w.measurements) < 3 {
		w.measurements = append(w.measurements, i)
		return
	}

	// If we have more than 3 measurements then "move" the window
	// and append the new measurement
	w.measurements = w.measurements[1:]

	w.measurements = append(w.measurements, i)
}

func (w *window) Sum() int {
	// Calculate the sum for the current window
	sum := 0
	for _, m := range w.measurements {
		sum += m
	}

	return sum

}

func (w *window) Compare() {
	// If the previousWindow value is populated compare the current window
	// sum against the previous window sum
	if w.previousWindow != 0 {
		if w.Sum() > w.previousWindow {
			w.increase++
		}
	}

	// Update the previous window sum to equal the current sum as the
	// window will slide
	w.previousWindow = w.Sum()
}

// A second implementation inspired by
func main(f string) int {
	scanner, close, err := fileReader(f)
	defer close()
	if err != nil {
		fmt.Printf("could not open file, got err: %s", err)
	}

	w := &window{}
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("could not parse: %s to int, got err: %s\n", scanner.Text(), err)
		}

		w.Populate(measurement)
		if len(w.measurements) < 3 {
			continue
		}

		w.Compare()
	}

	return w.increase
}

// Imagine that we need to write a function that compares a sliding average of
// measurements against one another.
//
// The size of each window will be 3 measurements, and we wish to record the
// number of times the sum of measurements in each sliding window increases
// compared to the last.
//
// Considering the above example:
//
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
//
// Start by comparing the first and second three-measurement windows. The
// measurements in the first window are marked A (199, 200, 208); their sum is
// 199 + 200 + 208 = 607. The second window is marked B (200, 208, 210); its sum
// is 618. The sum of measurements in the second window is larger than the sum
// of the first, so this first comparison increased.
//
// Your goal now is to count the number of times the sum of measurements in this
// sliding window increases from the previous sum. So, compare A with B, then
// compare B with C, then C with D, and so on. Stop when there aren't enough
// measurements left to create a new three-measurement sum.
//
// In the above example, the sum of each three-measurement window is as follows:
//
// A: 607 (N/A - no previous sum)
// B: 618 (increased)
// C: 618 (no change)
// D: 617 (decreased)
// E: 647 (increased)
// F: 716 (increased)
// G: 769 (increased)
// H: 792 (increased)
//
// In this example, there are 5 sums that are larger than the previous sum.
//
// Consider sums of a three-measurement sliding window. How many sums are larger than the previous sum?
