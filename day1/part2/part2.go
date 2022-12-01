package main

import (
	"amin/adv2022/helperUtils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := helperUtils.ReadFileLines("./part2.txt")
	// Creates a SLice of type int64 and size 1.
	// TODO: difference between sLice and array and int64 and int
	caloriesPerElf := make([]int64, 1)
	var currentIndex int = 0
	for _, line := range lines {
		if line == "" {
			currentIndex += 1
			// you need to use the value of append
			// TODO: difference between = and =:
			caloriesPerElf = append(caloriesPerElf, 0)
		} else {
			// Convert the string to an int value, first trim the new line character
			n, err := strconv.ParseInt(strings.TrimRight(line, "\r\n"), 10, 0)
			if err != nil {
				fmt.Println(err)
				return
			}

			caloriesPerElf[currentIndex] += n
		}
	}

	// Sort the array in descending order
	sort.Slice(caloriesPerElf, func(i, j int) bool {
		return caloriesPerElf[i] > caloriesPerElf[j]
	})

	// Get the first 3 elements from the sorted array
	biggest := caloriesPerElf[:3]

	var total int64 = 0
	// Iterate over the array and add each element to the total
	for _, n := range biggest {
		total += n
	}

	fmt.Println(total)
}
