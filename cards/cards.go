package cards

import "finances/entrypoint"

type Card interface {
	DoTheNeedFul(config entrypoint.Config)
}
