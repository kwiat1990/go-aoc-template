package main

import (
	"os"

	"go-aoc-template/internal"
)

func main() {
	if len(os.Args) > 1 {
		internal.HandleCommands()
	} else {
		internal.RunAllSolutions()
	}
}
