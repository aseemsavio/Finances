package utils

import (
	"fmt"
	"os"
)

// ReadFile looks for a file with a given file name and returns it.
func ReadFile(path string) *os.File {
	file, error := os.Open(path)
	if error != nil {
		fmt.Println("Error occurred while reading the Csv:", error)
	}
	return file
}
