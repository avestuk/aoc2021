package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Gamma rate is the most popular bit from left to right
// Epsilon rate is the least popular bit from left to right

// So we need to iterate over each input and count the values of each bit
// At the end of the input we have the most popular values for each position
// This is the gamma rate
// The epsilon rate will be the opposite

func TestParseInput(t *testing.T) {

	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	wantOxygen := "10111"
	wantCO := "01010"

	gotOxygen := ParseInput(input, 0, true)
	gotCO := ParseInput(input, 0, false)

	if gotOxygen != wantOxygen {
		t.Errorf("gotOxygen: %s != wantOxygen: %s", gotOxygen, wantOxygen)
	}
	if gotCO != wantCO {
		t.Errorf("gotCO: %s != wantCO: %s", gotCO, wantCO)
	}
}

func TestBinaryStringToInt(t *testing.T) {
	got, err := BinaryStringToInt("10110")
	if err != nil {
		t.Errorf("got err: %s", err)
	}

	want := int64(22)

	if got != want {
		t.Errorf("got: %d != %d", got, want)
	}
}

func TestDay3P1(t *testing.T) {
	got := Day3P1("./day3_test.txt")

	want := int64(230)

	if got != want {
		t.Errorf("got: %d != %d", got, want)
	}
}

func TestReturnRune(t *testing.T) {
	tests := []struct {
		name       string
		m          map[rune]int
		returnHigh bool
		want       rune
	}{
		{
			name:       "return one",
			m:          map[rune]int{'0': 10, '1': 100},
			returnHigh: true,
			want:       '1',
		},
		{
			name:       "tie break one",
			m:          map[rune]int{'0': 10, '1': 10},
			returnHigh: true,
			want:       '1',
		},
		{
			name: "return zero",
			m:    map[rune]int{'0': 10, '1': 100},
			want: '0',
		},
		{
			name: "tie break zero",
			m:    map[rune]int{'0': 10, '1': 10},
			want: '0',
		},
	}

	for _, test := range tests {

		gotRune := returnRune(test.m, test.returnHigh)

		if gotRune != test.want {
			t.Errorf("%s: %c != %c", test.name, gotRune, test.want)
		}
	}

}

func TestFilter(t *testing.T) {

	input := []string{
		"11011",
		"11011",
		"10011",
		"10011",
		"11011",
	}

	got := filter(input, 1, '0')

	require.Len(t, got, 2)
	require.Equal(t, "10011", got[0])

}
