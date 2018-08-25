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
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("Start of graphme.go.\n")

	//detect if our current dir is in a terraform project
	//grab our path
	curdir, err := os.Getwd()
	check(err)
	
	// fmt.Print(curdir)
	path := filepath.Join(curdir, ".terraform")
	
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0770)
	}

	//call `terraform init` before calling `terraform graph` (required by terraform)
	terraformInit := exec.Command("cmd", "/C", "terraform", "init")
	terraformInit.Output()

	//call `terraform graph` to generate a DOT file
	terrformGraph := exec.Command("cmd", "/C", "terraform", "graph", ">", "./.terraform/terraformgraph.dot")
	terrformGraph.Output()

	//process the DOT file
	dat, err := ioutil.ReadFile("./.terraform/terraformgraph.dot")
	check(err)

	//scrub unneeded items
	lines := strings.Split(string(dat), "\n")
	for i, line := range lines {

			//scrub
			if strings.Contains(line, "boundary") {
					lines[i] = ""
			}
			if strings.Contains(line, "[root] root") {
				lines[i] = ""
			}
			if strings.Contains(line, "[root] provider") {
				lines[i] = ""
			}

			//nodes
			if strings.Contains(line, "aws_vpc.") && strings.Contains(line, "[label") {
				fmt.Print("Found vpc node\n", line, "\n")
				fmt.Print("Update vpc node\n", strings.Replace(line, "box", "ellipse", -1), "\n")
				lines[i] = strings.Replace(line, "box", "ellipse", -1)
			}
			if strings.Contains(line, "aws_subnet.") && strings.Contains(line, "[label") {
				fmt.Print("Found subnet node\n", line, "\n")
			}
			
			//relations/edges
			if strings.Contains(line, "aws_subnet.") && strings.Contains(line, "aws_vpc.") {
				fmt.Print("Found vpc to subnet relation\n", line, "\n")
			}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./.terraform/scrub.dot", []byte(output), 0644)
	check(err)

	//parse dot file with graphviz parser and modify
	scrubdat, err := ioutil.ReadFile("./.terraform/scrub.dot")
	check(err)

	graph, err := gographviz.Read(scrubdat)
	check(err)

	graph.SetName("tfgraph")

	//add node
//	graph.AddNode("","bob",nil)


	s := graph.String()	
	
	//output modified dot file
	ioutil.WriteFile("./.terraform/tfgraph.dot", []byte(s), 0644)

	//call graphviz to convert the DOT file to SVG
	gvConvert := exec.Command("cmd", "/C", "dot", "-Tsvg", "./.terraform/tfgraph.dot",  ">", "./graph.svg")
	gvConvert.Output()

	//TODO add this as an optional flag
	//open svg file
	openSvg := exec.Command("powershell.exe", "-Command", "./graph.svg")
	openSvg.Output()

	fmt.Printf("End of graphme.go.\n")
}

//streamline error checking
func check(e error) {
    if e != nil {
        panic(e)
    }
}