package helperUtils

import (
	"bufio" // Import the "bufio" package for reading files line by line.
	"fmt"   // Import the "fmt" package for logging errors.
	"os"    // Import the "os" package for working with files.
)

// Define the function
func ReadFileLines(filePath string) []string {
	// Open the file at the given file path.
	file, err := os.Open(filePath)
	// If there is an error opening the file, log the error and return an empty array of strings.
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	// Defer closing the file until the function returns.
	defer file.Close()

	// Create a new scanner for reading the file line by line.
	scanner := bufio.NewScanner(file)

	// Create an array to store the lines of the file.
	var lines []string

	// Use a for loop to iterate over the lines in the file.
	for scanner.Scan() {
		// Append the current line to the array of lines.
		lines = append(lines, scanner.Text())
	}

	// If there is an error while scanning the file, log the error.
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Return the array of lines.
	return lines
}
