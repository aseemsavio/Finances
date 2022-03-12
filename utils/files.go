package utils

import (
	"fmt"
	"os"
)

// ReadFile
// Reads and returns a file object.
func ReadFile(path string) *os.File {
	file, error := os.Open(path)
	if error != nil {
		fmt.Println("Error occurred while reading the Csv:", error)
	}
	return file
}
