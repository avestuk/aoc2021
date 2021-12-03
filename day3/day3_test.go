package day3

import "testing"

// Gamma rate is the most popular bit from left to right
// Epsilon rate is the least popular bit from left to right

// So we need to iterate over each input and count the values of each bit
// At the end of the input we have the most popular values for each position
// This is the gamma rate
// The epsilon rate will be the opposite

func TestParseInput(t *testing.T) {

	tests := map[string]struct {
		Input       []string
		wantGamma   string
		wantEpsilon string
	}{
		"simple case": {

			Input:       []string{"00100"},
			wantGamma:   "00100",
			wantEpsilon: "11011",
		},
		"Double input": {

			Input:       []string{"00100", "11011", "11011"},
			wantGamma:   "11011",
			wantEpsilon: "00100",
		},
	}

	for name, test := range tests {
		gotGamma, gotEpsilon := ParseInput(test.Input)

		if gotGamma != test.wantGamma {
			t.Errorf("%s: gotGamma: %s != wantGamma: %s", name, gotGamma, test.wantGamma)
		}
		if gotEpsilon != test.wantEpsilon {
			t.Errorf("%s: gotEpsilon: %s != wantEpsilon: %s", name, gotEpsilon, test.wantEpsilon)
		}

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

	want := int64(198)

	if got != want {
		t.Errorf("got: %d != %d", got, want)
	}
}
