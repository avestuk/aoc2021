package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Read input
func TestDrawNumbers(t *testing.T) {
	b := &Bingo{
		drawnIndex: 0,
		numbers:    []string{"7", "4", "9", "5", "11", "17", "23", "2", "0", "14", "21"},
	}

	want := []string{"7", "4", "9", "5", "11"}

	var got []string
	for i := 0; i < 5; i++ {
		n, err := b.DrawNumbers()
		require.NoError(t, err)
		got = append(got, n)
	}
	require.Equal(t, want, got)

	want = []string{"17", "23", "2", "0", "14", "21"}
	var got2 []string
	for i := 0; i < 6; i++ {
		n, err := b.DrawNumbers()
		require.NoError(t, err)
		got2 = append(got2, n)
	}

	require.Equal(t, want, got2)
}

func TestCheckHorizontal(t *testing.T) {
	numbers := []string{"22", "13", "17", "11", "0"}
	wantBoard := Board{
		[]string{"22", "13", "17", "11", "0"},
		[]string{"8", "2", "23", "4", "4"},
		[]string{"21", " 9", "14", "16", "7"},
		[]string{"6", "10", "3", "18", "5"},
		[]string{"1", "12", "20", "15", "9"},
	}

	require.True(t, CheckHorizontal(wantBoard, numbers))

	numbers = []string{"17", "23", "14", "3", "20"}
	require.True(t, CheckHorizontal(wantBoard, numbers))
}

func TestVertical(t *testing.T) {
	numbers := []string{"17", "23", "14", "3", "20"}
	wantBoard := Board{
		[]string{"22", "13", "17", "11", "0"},
		[]string{"8", "2", "23", "4", "4"},
		[]string{"21", "9", "14", "16", "7"},
		[]string{"6", "10", "3", "18", "5"},
		[]string{"1", "12", "20", "15", "9"},
	}

	require.True(t, CheckVertical(wantBoard, numbers, []int{2}))
}

func TestParseInput(t *testing.T) {

	input, boards, err := ParseInput("./day4_test.txt")
	require.NoError(t, err)

	require.Len(t, input, 27)
	wantBoards := Boards{
		Board{
			[]string{"22", "13", "17", "11", "0"},
			[]string{"8", "2", "23", "4", "24"},
			[]string{"21", "9", "14", "16", "7"},
			[]string{"6", "10", "3", "18", "5"},
			[]string{"1", "12", "20", "15", "19"},
		},
		{
			[]string{"3", "15", "0", "2", "22"},
			[]string{"9", "18", "13", "17", "5"},
			[]string{"19", "8", "7", "25", "23"},
			[]string{"20", "11", "10", "24", "4"},
			[]string{"14", "21", "16", "12", "6"},
		},
		{
			[]string{"14", "21", "17", "24", "4"},
			[]string{"10", "16", "15", "9", "19"},
			[]string{"18", "8", "23", "26", "20"},
			[]string{"22", "11", "13", "6", "5"},
			[]string{"2", "0", "12", "3", "7"},
		},
	}

	for i, gb := range boards {
		for j, row := range gb {
			require.Equal(t, wantBoards[i][j], row)
		}
	}

	require.Equal(t, wantBoards, boards)
}

func TestCalculateScore(t *testing.T) {
	input, _, err := ParseInput("./day4_test.txt")
	require.NoError(t, err)

	b := &Bingo{
		drawnIndex: 12,
		numbers:    input,
	}

	board := Board{
		[]string{"14", "21", "17", "24", "4"},
		[]string{"10", "16", "15", "9", "19"},
		[]string{"18", "8", "23", "26", "20"},
		[]string{"22", "11", "13", "6", "5"},
		[]string{"2", "0", "12", "3", "7"},
	}

	wantScore := 4512

	gotScore, err := b.CalculateScore(board)
	require.NoError(t, err)

	require.Equal(t, wantScore, gotScore)

}

func TestDay4P1(t *testing.T) {
	gotScore, err := Day4P1("./day4_test.txt")
	require.NoError(t, err)
	require.Equal(t, 4512, gotScore)

	// Sum values
	// []int len: 13, cap: 16, [10,16,15,19,18,8,26,20,22,13,6,12,3]
}

func TestDay4P2(t *testing.T) {
	gotScore, err := Day4P2("./day4_test.txt")
	require.NoError(t, err)
	require.Equal(t, 1924, gotScore)

	// Sum values
	// []int len: 13, cap: 16, [10,16,15,19,18,8,26,20,22,13,6,12,3]
}

// Continue
// When a winner is found find all the "unmarked" numbers on the winning board(s)

// --- Day 4: Giant Squid ---
//
// You're already almost 1.5km (almost a mile) below the surface of the ocean,
// already so deep that you can't see any sunlight. What you can see, however,
// is a giant squid that has attached itself to the outside of your submarine.
//
// Maybe it wants to play bingo?
//
// Bingo is played on a set of boards each consisting of a 5x5 grid of numbers.
// Numbers are chosen at random, and the chosen number is marked on all boards
// on which it appears. (Numbers may not appear on all boards.) If all numbers
// in any row or any column of a board are marked, that board wins. (Diagonals
// don't count.)
//
// The submarine has a bingo subsystem to help passengers (currently, you and
// the giant squid) pass the time. It automatically generates a random order in
// which to draw numbers and a random set of boards (your puzzle input). For
// example:
//
// 7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1
//
// 22 13 17 11  0
//  8  2 23  4 24
// 21  9 14 16  7
//  6 10  3 18  5
//  1 12 20 15 19
//
//  3 15  0  2 22
//  9 18 13 17  5
// 19  8  7 25 23
// 20 11 10 24  4
// 14 21 16 12  6
//
// 14 21 17 24  4
// 10 16 15  9 19
// 18  8 23 26 20
// 22 11 13  6  5
//  2  0 12  3  7
//
// After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):
//
// 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
//  8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
// 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
//  6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
//  1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
//
// After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:
//
// 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
//  8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
// 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
//  6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
//  1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
//
// Finally, 24 is drawn:
//
// 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
//  8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
// 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
//  6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
//  1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
//
// At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).
//
// The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.
//
// To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?
