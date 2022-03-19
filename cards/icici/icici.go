package icici

import (
	"finances/entrypoint"
	"finances/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

type Record struct {
	TransactionDate string
	Details         string
	Amount          int64
	ReferenceNumber int64
}

func DoTheNeedful(config entrypoint.Config) {
	file := utils.ReadFile(config.File)
	recordsArray := utils.Csv(file)
	transactionalDetails := getTransactionDetails(recordsArray)
	correctAnomalies(transactionalDetails)
	debited, credited := getDebitedAndCreditArrays(transactionalDetails)
	displayTable(debited)
	displayTable(credited)
}

// getDebitedAndCreditArrays returns two slices of debited and credited records.
func getDebitedAndCreditArrays(lines [][]string) ([][]string, [][]string) {
	var debited [][]string
	var credited [][]string
	for _, line := range lines {
		if strings.HasSuffix(strings.TrimSpace(line[2]), "Dr.") {
			line[2] = strings.TrimSuffix(line[2], "Dr.")
			debited = append(debited, line)
		}
		if strings.HasSuffix(strings.TrimSpace(line[2]), "Cr.") {
			line[2] = strings.TrimSuffix(line[2], "Cr.")
			credited = append(credited, line)
		}
	}
	return debited, credited
}

func correctAnomalies(details [][]string) {
	for i := range details {
		if i > 0 {
			for j := range details[i] {
				switch j {
				case 0:
					date := details[i][j]
					details[i][j] = strings.Replace(date, ",", "-", -1)
				case 2:
					malformedAmountString := details[i][j]
					correctedAmount := strings.Replace(malformedAmountString, ",", "", -1)
					details[i][j] = correctedAmount
				}
			}
		}
	}
}

func displayTable(transactionalDetails [][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Transaction Date",
		"Details",
		"Amount (INR)",
		"Reference Number",
	})
	for i := range transactionalDetails {
		t.AppendRow(table.Row{
			transactionalDetails[i][0],
			transactionalDetails[i][1],
			transactionalDetails[i][2],
			transactionalDetails[i][3],
		})
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
