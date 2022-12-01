package main

import (
	"amin/adv2022/helperUtils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Call the function and print the result
	lines := helperUtils.ReadFileLines("./part1.txt")
	var maxCalories int64 = 0
	var currentCalories int64 = 0
	for _, line := range lines {
		if line == "" {
			currentCalories = 0
		} else {
			// Convert the string to an int value, first trim the new line character
			n, err := strconv.ParseInt(strings.TrimRight(line, "\r\n"), 10, 0)
			if err != nil {
				fmt.Println(err)
				return
			}

			currentCalories += n

			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
		}
	}

	fmt.Println(maxCalories)
}
