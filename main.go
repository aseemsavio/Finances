package main

import (
	"finances/core"
	"finances/utils"
)

func main() {
	file := utils.LocalFile(utils.ReadFile("finance.csv"))
	records := file.Csv()
	core.GetMonthsSpending(records)
}
