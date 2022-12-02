package main

import (
	"amin/adv2022/helperUtils"
	"fmt"
	"strings"
)

func main() {
	// Call the function and print the result
	lines := helperUtils.ReadFileLines("./part2.txt")

	var shapes = map[string]int{
		"X":   0, // lose
		"Y":   3, // draw
		"Z":   6, // win
		"A X": 3,
		"A Y": 1,
		"A Z": 2,
		"B X": 1,
		"B Y": 2,
		"B Z": 3,
		"C X": 2,
		"C Y": 3,
		"C Z": 1,
	}

	count := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		myTurn := strings.Fields(trimmed)[1]

		count += shapes[myTurn]
		count += shapes[trimmed]
	}

	fmt.Println(count)
}
