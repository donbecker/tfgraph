// Example Usage
// normal command line (without this file)
// `terraform graph | dot -Tsvg > graph.svg`
// what this file should do
// `terraform graph | graphprocess.go | dot -Tsvg > graph.svg`

// Execution from vscode console
// go run graphme.go graphprocess.go

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Start of graphme.go.\n")

	ProcessDOT()

	fmt.Printf("End of graphme.go.\n")
}