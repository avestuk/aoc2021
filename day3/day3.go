package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day3P1(f string) int64 {
	scanner, close, err := fileReader(f)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer close()

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	g, e := ParseInput(input)

	gamma, err := BinaryStringToInt(g)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	epsilon, err := BinaryStringToInt(e)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return gamma * epsilon
}

func ParseInput(input []string) (string, string) {

	popularity := make([]map[rune]int, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		popularity[i] = map[rune]int{}
	}

	//popularity := []map[rune]int{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}
	//popularity := make([]map[rune]int, 2, 10)

	for _, s := range input {
		for i, c := range s {
			popularity[i][c] += 1
		}
	}

	var (
		g string
		e string
	)
	for _, popMap := range popularity {
		if popMap['0'] > popMap['1'] {
			g = g + "0"
			e = e + "1"
		} else {
			g = g + "1"
			e = e + "0"
		}
	}

	return g, e
}

func BinaryStringToInt(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}

func fileReader(f string) (*bufio.Scanner, func() error, error) {
	// Handle reading from a file
	file, err := os.Open(f)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open file, got err: %w", err)
	}

	return bufio.NewScanner(file), file.Close, err
}
