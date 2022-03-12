package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

//Csv creates CSVs out of a file
func Csv(file *os.File) [][]string {
	reader := csv.NewReader(file)
	lines, error := reader.ReadAll()
	if error != nil {
		fmt.Println("Error occurred while parsing the CSV file:", error)
	}
	return lines
}
