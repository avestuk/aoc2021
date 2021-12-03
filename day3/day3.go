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

	o := ParseInput(input, 0, true)
	c := ParseInput(input, 0, false)

	oxygen, err := BinaryStringToInt(o)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	co2, err := BinaryStringToInt(c)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return oxygen * co2
}

func ParseInput(input []string, byteCounter int, findOxy bool) string {
	// Loop over our input
	// For each first byte determine the least and most popular
	popularity := make(map[rune]int)
	for _, s := range input {
		popularity[[]rune(s)[byteCounter]] += 1
	}

	// Determine the most popular value
	mostPopular := returnRune(popularity, findOxy)

	// Filter
	sived := filter(input, byteCounter, mostPopular)

	if len(sived) > 1 {
		byteCounter += 1
		return ParseInput(sived, byteCounter, findOxy)
	}

	return sived[0]
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

func returnRune(popularity map[rune]int, returnHigh bool) rune {
	zeroCount := popularity['0']
	oneCount := popularity['1']

	if returnHigh {
		if oneCount >= zeroCount {
			return '1'
		}
		return '0'
	}
	if zeroCount <= oneCount {
		return '0'
	}
	return '1'
}

func filter(input []string, byteCounter int, mostPopular rune) []string {
	sived := []string{}
	for _, s := range input {
		if []rune(s)[byteCounter] == mostPopular {
			sived = append(sived, s)
		}
	}

	return sived
}
