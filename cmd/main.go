package main

import (
	"fmt"

	"github.com/avestuk/aoc2021/day2"
)

func main() {
	const (
		dayOneInput string = "./day1/day1.txt"
		dayTwoInput string = "./day2/day2.txt"
	)

	//count, err := day1.Day1P1(dayOneInput)
	//count, err := day1.Day1P22(dayOneInput)
	count := day2.Day2P1(dayTwoInput)

	fmt.Println("count was: ", count)
}
