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
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Printf("Start of graphme.go.\n")

	//detect if our current dir is in a terraform project
	//grab our path
	curdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Print(curdir)


	//call `terraform graph` to generate a DOT file
	tfCmd := exec.Command("cmd", "/C", "terraform", "graph", ">", "graph.dot")
	tfCmd.Output()

	//proceess the DOT file
	ProcessDOT()

	//call graphviz to convert the DOT file to SVG
	gvCmd := exec.Command("cmd", "/C", "dot", "-Tsvg", "graph.dot",  ">", "tfgraph.svg")
	gvCmd.Output()

	fmt.Printf("End of graphme.go.\n")
}