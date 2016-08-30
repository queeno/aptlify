package main

import (
	"github.com/queeno/aptlify/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Run(cmd.RootCommand(), os.Args[1:], true))
}
