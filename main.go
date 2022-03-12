package main

import (
	"finances/core"
	"finances/entrypoint"
)

func main() {
	config := entrypoint.CompileArguments()
	core.DoTheNeedful(config)
}
