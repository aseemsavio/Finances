package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

type File struct {
	*os.File
}

func LocalFile(file *os.File) File {
	return File{file}
}

//Csv
// Creates CSV out of a file
func (file *File) Csv() [][]string {
	reader := csv.NewReader(file)
	lines, error := reader.ReadAll()
	if error != nil {
		fmt.Println("Error occurred while parsing the CSV file:", error)
	}
	return lines
}
