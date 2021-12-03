package main

import (
	"fmt"

	"github.com/avestuk/aoc2021/day3"
)

func main() {
	const (
		dayOneInput   string = "./day1/day1.txt"
		dayTwoInput   string = "./day2/day2.txt"
		dayThreeInput string = "./day3/day3.txt"
	)

	//count, err := day1.Day1P1(dayOneInput)
	//count, err := day1.Day1P22(dayOneInput)
	count := day3.Day3P1(dayThreeInput)
	fmt.Println("count was: ", count)
}
