package helperUtils

import (
	"bufio"
	"fmt"
	"os"
)

// Define the function
func ReadFileLines(filePath string) []string {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	defer file.Close()

	// Read the file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Return the array of lines
	return lines
}
