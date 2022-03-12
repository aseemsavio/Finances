package core

func GetMonthsSpending(lines [][]string) {
	for i := range lines {
		for j := range lines[i] {
			print(lines[i][j] + "--")
		}
		println()
	}
}
