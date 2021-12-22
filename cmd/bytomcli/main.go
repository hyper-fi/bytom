package main

import (
	"runtime"

	cmd "github.com/hyper-fi/bytom/cmd/bytomcli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
