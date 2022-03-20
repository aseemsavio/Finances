package icici

import (
	"finances/entrypoint"
	"finances/utils"
	"finances/utils/database"
	"fmt"
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

const (
	CardName = "ICICI-AMAZON-Credit-CARD"
	Credit   = "CREDIT"
	Debit    = "DEBIT"
)

func DoTheNeedful(config entrypoint.Config) {
	file := utils.ReadFile(config.File)
	recordsArray := utils.Csv(file)
	transactionalDetails := getTransactionDetails(recordsArray)
	correctAnomalies(transactionalDetails)
	debited, credited := getDebitAndCreditArrays(transactionalDetails)
	displayTable(debited)
	displayTable(credited)

	db, _ := database.NewSpreadsheetService()
	/*	data, _ := db.GetData(config.SpreadSheetId, "Sheet1!A1:D3")

		for _, row := range data.ValueRange.Values {
			fmt.Printf("%s, %s, %s, %s\n", row[0], row[1], row[2], row[3])
		}
	*/
	for _, value := range debited {
		db.PutData(config.SpreadSheetId, "A1", []interface{}{value[0], value[1], value[2], value[3], CardName, Debit})
	}

	for _, value := range credited {
		db.PutData(config.SpreadSheetId, "A1", []interface{}{value[0], value[1], value[2], value[3], CardName, Credit})
	}

	fmt.Printf("Service: %+v", db)
}

// getDebitAndCreditArrays returns two slices of debited and credited records.
func getDebitAndCreditArrays(lines [][]string) ([][]string, [][]string) {
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
				/* Date format is weird (with commas!!) in the CSV. This fixes it. */
				case 0:
					date := details[i][j]
					details[i][j] = strings.Replace(date, ",", "-", -1)
				/* Gets rid of the commas in the amount */
				case 2:
					malformedAmountString := details[i][j]
					moneyMoneyMoney := strings.Replace(malformedAmountString, ",", "", -1)
					details[i][j] = moneyMoneyMoney
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
			/* This is where the transaction details start in the CSV file.
			Hence, anything above this is ignored.*/
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

// toInterfaceSlice converts a slice of strings into a slice of interface.
func toInterfaceSlice(stringSlice [][]string) [][]interface{} {
	myInterface := make([][]interface{}, len(stringSlice))
	for i, _ := range myInterface {
		myInterface[i] = make([]interface{}, len(stringSlice[0]))
	}

	for i, row := range stringSlice {
		for j, value := range row {
			myInterface[i][j] = value
		}
	}
	return myInterface
}
