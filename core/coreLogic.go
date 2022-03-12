package core

func GetMonthsSpending(lines [][]string) {
	transactionalDetails := getTransactionDetails(lines)
	for i := range transactionalDetails {
		for j := range transactionalDetails[i] {
			print(transactionalDetails[i][j], " ")
		}
		println()
	}
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
