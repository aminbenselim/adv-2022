package main

import (
	"amin/adv2022/helperUtils"
	"fmt"
	"strings"
)

func main() {
	// Call the function and print the result
	lines := helperUtils.ReadFileLines("./part1.txt")

	var shapes = map[string]int{
		"X":   1, // Rock
		"Y":   2, // Paper
		"Z":   3, // Sicssors
		"A X": 3,
		"A Y": 6,
		"A Z": 0,
		"B X": 0,
		"B Y": 3,
		"B Z": 6,
		"C X": 6,
		"C Y": 0,
		"C Z": 3,
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
