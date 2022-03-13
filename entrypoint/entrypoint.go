package entrypoint

import "flag"

type Config struct {
	Card string
	File string
}

// CompileArguments compiles the command line argument and outputs a Config object.
func CompileArguments() Config {
	card := flag.String("card", "icici", "Usage: -card=icici")
	file := flag.String("file", "finance.csv", "Usage: -File=finance.csv")
	flag.Parse()
	return Config{
		Card: *card,
		File: *file,
	}
}
