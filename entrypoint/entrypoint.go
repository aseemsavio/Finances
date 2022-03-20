package entrypoint

import "flag"

type Config struct {
	Card          string
	File          string
	SpreadSheetId string
}

// CompileArguments compiles the command line argument and outputs a Config object.
func CompileArguments() Config {
	card := flag.String("card", "icici", "Usage: -card=icici")
	file := flag.String("file", "finance.csv", "Usage: -File=finance.csv")
	spreadSheetId := flag.String("spreadsheetId", "", "")
	flag.Parse()
	return Config{
		Card:          *card,
		File:          *file,
		SpreadSheetId: *spreadSheetId,
	}
}
