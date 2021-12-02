package main

import (
	"fmt"

	"github.com/avestuk/aoc2021/day1"
)

func main() {
	const (
		dayOneInput string = "./day1/day1.txt"
	)

	//count, err := day1.Day1P1(dayOneInput)
	count, err := day1.Day1P22(dayOneInput)
	if err != nil {
		fmt.Printf("got err: %s\n", err)
		return
	}

	fmt.Println("count was: ", count)
}
