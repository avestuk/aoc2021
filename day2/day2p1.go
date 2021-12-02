package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sub struct {
	horizontal int
	depth      int
	aim        int
}

func (s *sub) Forward(i int) {
	s.horizontal += i
	s.depth += s.aim * i
}

func (s *sub) Down(i int) {
	s.aim += i
}

func (s *sub) Up(i int) {
	s.aim -= i
}

func (s *sub) Parse(cmd string) error {
	// Split the command into the movement and units
	movement := strings.Split(cmd, " ")[0]
	units := strings.Split(cmd, " ")[1]

	unit, err := strconv.Atoi(units)
	if err != nil {
		return fmt.Errorf("could not convert: %s to int, got err: %w", units, err)
	}

	switch movement {
	case "up":
		s.Up(unit)
	case "down":
		s.Down(unit)
	case "forward":
		s.Forward(unit)
	}

	return nil
}

func (s *sub) Multiply() int {
	return s.depth * s.horizontal
}

func Day2P1(f string) int {

	scanner, close, err := fileReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer close()

	s := &sub{}
	for scanner.Scan() {
		cmd := scanner.Text()

		if err := s.Parse(cmd); err != nil {
			fmt.Println(err)
		}
	}

	return s.Multiply()
}

func fileReader(f string) (*bufio.Scanner, func() error, error) {
	// Handle reading from a file
	file, err := os.Open(f)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open file, got err: %w", err)
	}

	return bufio.NewScanner(file), file.Close, err
}
