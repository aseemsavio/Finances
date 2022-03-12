package icici

import (
	"finances/entrypoint"
	"finances/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func DoTheNeedful(config entrypoint.Config) {
	file := utils.ReadFile(config.File)
	records := utils.Csv(file)
	getMonthsSpending(records)
}

func getMonthsSpending(lines [][]string) {
	transactionalDetails := getTransactionDetails(lines)
	displayTable(transactionalDetails)
}

func displayTable(transactionalDetails [][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		transactionalDetails[0][0],
		transactionalDetails[0][1],
		transactionalDetails[0][2],
		transactionalDetails[0][3],
	})
	for i := range transactionalDetails {
		if i > 0 {
			t.AppendRow(table.Row{
				transactionalDetails[i][0],
				transactionalDetails[i][1],
				transactionalDetails[i][2],
				transactionalDetails[i][3],
			})
		}
	}
	t.AppendSeparator()
	t.Render()
}

func getTransactionDetails(lines [][]string) [][]string {
	var transactionDetails [][]string
	start := false
	for i := range lines {
		var row []string
		for j := range lines[i] {
			value := lines[i][j]
			if value == "Transaction Details" {
				start = true
				break
			}
			if start {
				if value != "" {
					row = append(row, value)
				}
			}
		}
		if len(row) > 0 {
			transactionDetails = append(transactionDetails, row)
			row = []string{}
		}
	}
	return transactionDetails
}
