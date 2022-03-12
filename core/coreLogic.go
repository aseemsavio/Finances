package core

import (
	"finances/cards/icici"
	"finances/entrypoint"
)

func DoTheNeedful(config entrypoint.Config) {
	switch config.Card {
	case "icici":
		icici.DoTheNeedful(config)
	case "hdfc":
	}
}
