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
	"github.com/awalterschulze/gographviz"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	fmt.Printf("Start of graphme.go.\n")

	//detect if our current dir is in a terraform project
	//grab our path
	curdir, err := os.Getwd()
	check(err)
	
	fmt.Print(curdir)

	//call `terraform graph` to generate a DOT file
	tfCmd := exec.Command("cmd", "/C", "terraform", "graph", ">", "graph.dot")
	tfCmd.Output()

	//process the DOT file
	dat, err := ioutil.ReadFile("./graph.dot")
	check(err)

	graph, err := gographviz.Read(dat)
	check(err)

	graph.SetName("tfgraph")
	s := graph.String()	
	
	//output modified dot file
	ioutil.WriteFile("tfgraph.dot", []byte(s), 0644)

	//call graphviz to convert the DOT file to SVG
	gvCmd := exec.Command("cmd", "/C", "dot", "-Tsvg", "tfgraph.dot",  ">", "tfgraph.svg")
	gvCmd.Output()

	fmt.Printf("End of graphme.go.\n")
}

//streamline error checking
func check(e error) {
    if e != nil {
        panic(e)
    }
}