package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bingo struct {
	numbers    []string
	drawnIndex int
}

type Board [][]string

type Boards []Board

func (b *Bingo) DrawNumbers() (string, error) {
	if len(b.numbers) == b.drawnIndex {
		return "", fmt.Errorf("numbers is: %d in length, index: %d", len(b.numbers), b.drawnIndex)
	}

	i := b.numbers[b.drawnIndex]
	b.drawnIndex += 1

	return i, nil
}

func (b *Bingo) CheckBoards(boards Boards) Board {
	numbers := b.numbers[:b.drawnIndex]

	for _, board := range boards {
		if boardMatch := CheckHorizontal(board, numbers); boardMatch {
			return board
		}
	}

	return Board{}
}

func CheckHorizontal(board Board, numbers []string) bool {
	for boardIndex, row := range board {
		var rowMatch []int
	Row:
		for rowIndex, bi := range row {
			for _, n := range numbers {
				if bi == n {
					rowMatch = append(rowMatch, rowIndex)
					// We have a winning board
					if len(rowMatch) == 5 {
						return true
					}
					continue Row
				}
			}

			// If there was a board match in the first row check
			// for vertical matches.
			if boardIndex == 0 && len(rowMatch) != 0 {
				if boardMatches := CheckVertical(board, numbers, rowMatch); boardMatches {
					return boardMatches
				}
			}
		}
	}
	return false
}

func CheckVertical(board Board, numbers []string, columnIndex []int) bool {
	var boardMatches bool
	for _, colIndex := range columnIndex {
	Row:
		for _, row := range board {
			for _, n := range numbers {
				if row[colIndex] == n {
					boardMatches = true
					continue Row
				}
			}
			return false
		}
		return boardMatches
	}
	return false
}

func (b *Bingo) CalculateScore(board Board) (int, error) {
	sum := 0
	var sumConstituents []int
	for _, row := range board {
	Row:
		for _, i := range row {
			for _, n := range b.numbers[:b.drawnIndex] {
				if i == n {
					continue Row
				}
			}

			j, err := strconv.Atoi(i)
			if err != nil {
				return 0, fmt.Errorf("could not convert: %s to int", i)
			}
			sumConstituents = append(sumConstituents, j)
			sum += j
		}
	}

	finalNumber, err := strconv.Atoi(b.numbers[b.drawnIndex-1])
	if err != nil {
		return 0, fmt.Errorf("could not convert: %s to int", b.numbers[b.drawnIndex])
	}

	return sum * finalNumber, nil
}

func Day4P1(f string) (int, error) {
	input, boards, err := ParseInput(f)
	if err != nil {
		return 0, err
	}

	b := &Bingo{
		drawnIndex: 0,
		numbers:    input,
	}

	for i := 0; i < len(b.numbers); i++ {
		b.DrawNumbers()
		// No point checking for winning boards until we've drawn 5 numbers
		if b.drawnIndex < 4 {
			continue
		}

		winningBoard := b.CheckBoards(boards)
		if len(winningBoard) == 0 {
			continue
		}

		return b.CalculateScore(winningBoard)
	}

	return 0, fmt.Errorf("did not find a winning board")
}

func Day4P2(f string) (int, error) {
	input, boards, err := ParseInput(f)
	if err != nil {
		return 0, err
	}

	b := &Bingo{
		drawnIndex: 0,
		numbers:    input,
	}

	boardCount := len(boards)

	var winningBoardsOrder 
	for i := 0; i < len(b.numbers); i++ {
		b.DrawNumbers()
		// No point checking for winning boards until we've drawn 5 numbers
		if b.drawnIndex < 4 {
			continue
		}

		wBoard := b.CheckBoards(boards)
		if len(wBoard) == 0 {
			continue
		}
		winningBoardsOrder[len(winningBoardsOrder)] = wBoard

		if len(winningBoardsOrder) == boardCount {
			return b.CalculateScore(winningBoardsOrder[boardCount-1])
		}
	}
	return 0, fmt.Errorf("something broke")
}

func fileReader(f string) (*bufio.Scanner, func() error, error) {
	// Handle reading from a file
	file, err := os.Open(f)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open file, got err: %w", err)
	}

	return bufio.NewScanner(file), file.Close, err
}

func ParseInput(f string) ([]string, Boards, error) {
	scanner, close, err := fileReader(f)
	if err != nil {
		return nil, nil, err
	}
	defer close()

	var (
		input  []string
		rows   [][]string
		boards Boards
	)
	for scanner.Scan() {
		// Handle empty lines
		if len(scanner.Text()) == 0 {
			continue
		}

		// Handle the numbers
		if len(input) == 0 {
			input = strings.Split(scanner.Text(), ",")
			continue
		}

		// Handle boards
		var row []string
		dirtyRow := strings.Split(scanner.Text(), " ")
		for _, entry := range dirtyRow {
			if entry != " " && entry != "" {
				row = append(row, entry)
			}
		}
		rows = append(rows, row)

		if len(rows) == 5 {
			boards = append(boards, rows)
			rows = [][]string{}
		}
	}

	return input, boards, nil
}

// Read input in sets of 5
// Check for complete rows of 5
// Continue
// When a winner is found find all the "unmarked" numbers on the winning board(s)
