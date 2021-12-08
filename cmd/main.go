package main

import (
	"fmt"

	"github.com/avestuk/aoc2021/day4"
)

func main() {
	const (
		dayOneInput   string = "./day1/day1.txt"
		dayTwoInput   string = "./day2/day2.txt"
		dayThreeInput string = "./day3/day3.txt"
		dayFourInput  string = "./day4/day4.txt"
	)

	//count, err := day1.Day1P1(dayOneInput)
	//count, err := day1.Day1P22(dayOneInput)
	count, err := day4.Day4P1(dayFourInput)
	if err != nil {
		fmt.Printf("got err: %s, err")
		return
	}
	fmt.Println("count was: ", count)
}
