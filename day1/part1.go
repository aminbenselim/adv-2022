package main

import (
	"fmt"

	"amin/adv2022/helperUtils"
)

func main() {
	// Call the function and print the result
	lines := helperUtils.ReadFileLines("./part1.txt")
	for _, line := range lines {
		fmt.Println(line)
	}
}
