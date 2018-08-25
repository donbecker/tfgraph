## SVG
* vscode: extensions: SVG Viewer v1.4.4 by cssho

## Icons
* AWS Icons Package
    * simple icons page: https://aws.amazon.com/architecture/icons/
    * direct link: https://s3-us-west-2.amazonaws.com/awswebanddesign/Architecture+Icons/AWS-Arch-Icon-Sets_Feb-18/PNG%2C+SVG%2C+EPS_18.02.22.zip
    * unzip
* PNG Icons used (in 'PNG, SVG, EPS' folder)
    * VPC
        * ..\General\General_virtualprivatecloud.png
        * ..\Compute\Compute_AmazonVPC.png
    * RDS 
        * ..\Database\Database_AmazonRDS.png
    * WAF
        * ..\Security Identity & Compliance\SecurityIdentityCompliance_AWSWAF.png
    * Route Table
    * Subnet
        * Missing
    * Security Group
    * EC2 Instance
        * ..\Compute\Compute_AmazonEC2_instance.png
    * Application Load Balancer
        * ..\Networking & Content Delivery\NetworkingContentDelivery_ElasticLoadBalancing_ApplicationLoadBalancer.png  
        * ..\Compute\Compute_ElasticLoadBalancing_ApplicationLoadBalancer.png
    * NAT Gateway
        * ..\Networking & Content Delivery\NetworkingContentDelivery_ElasticLoadBalancing_ApplicationLoadBalancer.png
        * ..\Compute\Compute_AmazonVPC_VPCNATgateway.png
    * Internet Gateway
        * ..\Compute\Compute_AmazonVPC_Internetgateway.png
    * VPC Peering
        * ..\Networking & Content Delivery\NetworkingContentDelivery_AmazonVPC_VPCpeering.png
        * ..\Compute\Compute_AmazonVPC_VPCpeering.png
    * Router
        * ..\Networking & Content Delivery\NetworkingContentDelivery_AmazonVPC_router.png

## Layout Style Notes

* Region: Outermost box, solid black line, transparent background, "AWS" logo in upper left, rounded corners, usually says "Region" at bottom. Not sure how multi-region diagrams look, initially only support one region.

* VPC: Second most outermost box, solid black line, transparent background, "VPC" logo in upper left, rounded corners.

* VPC Peering: Double-line joining the VPC boxes with VPC Peering icon in middle of line

* Availability Zone: Orange dashed line box with light semi-transparent orange background. Name (ex. 'Availability Zone 1') in orange centered left to right, above bottom line of AZ box. If stacked vertically, left-right edges extend outside VPC box. If stacked horizontally, top-bottom edges extend outside VPC box. 

* Subnet: Entirely within VPC box, lock icon in upper left, solid black line, transparent background. Public subnets have light green background, Private subnets have a light blue background. Subnet name and CIDR centered left to right, above bottom line.

* EC2 Instance: Orange box EC2 icon, Name in black text centered left to right and below box.

## Script notes

* The bulk of Terraform is written in Go. So it makes sense to do the same.
* Initial thought for distribution was to use a terraform module, but that would require user to manually invoke the embedded Go script. 
* Another option is writing a custom Terraform provider (plugin)
    * https://www.terraform.io/docs/extend/writing-custom-providers.html
    * https://www.terraform.io/docs/plugins/basics.html
    * To test your plugin, the easiest method is to copy your terraform binary to $GOPATH/bin and ensure that this copy is the one being used for testing. terraform init will search for plugins within the same directory as the terraform binary, and $GOPATH/bin is the directory into which go install will place the plugin executable.
    * Unfortunately it looks like retrieval of plugins is a bit of a mess
        * https://github.com/hashicorp/terraform/issues/15252
* So thinking we do this with a go script, distributed via a module.
    * User adds module ref
    * User runs `terraform init`
        * Downloads module, and untars
    * ??? where do we find the module under `.terraform` directory ???
    * User runs `graphme.go` script
* `terraform graph` scripts
    * located the rest of the terraform graph scripts
        * https://github.com/hashicorp/terraform/blob/master/command/graph.go
        * https://github.com/hashicorp/terraform/blob/master/terraform/graph.go
        * https://github.com/hashicorp/terraform/blob/master/dag/graph.go


* Yet another option is using the `local-exec` provisioner
    * https://www.terraform.io/docs/provisioners/local-exec.html
    * "The local-exec provisioner invokes a local executable after a resource is created. This invokes a process on the machine running Terraform, not on the resource."
    * So can we use this to call our Go script? 


* Scripts
    * `graphme.go`
        * Main (wrapper) script
        * "graphme.go" is dev name of this file, `tfgraph` will map to this file/be the release name
            * So for dev, cd into tfproject folder, run `terraform init` and then run `.\go run graphme.go`
        * Steps
            * Runs `terraform graph` which outputs DOT file
            * Calls `graphprocess.go` to process DOT file
            * Calls `graphviz` to render the SVG graph from DOT file

    * `graphprocess.go`
        * Updates the DOT file created by `terraform graph` command
        * Different rulesets used for different intended diagram types
        * Rulesets passed in as command line options or via defaults file
        * Sets the icon used on the nodes
        * Removes items from DOT file that are not desired
        * Looks like there is a DOT file parser for Go already
            * https://github.com/awalterschulze/gographviz
            * Example using this for Grails/Groovy
                * https://github.com/ilikeorangutans/grails-service-visualizer
                * https://ilikeorangutans.github.io/2014/05/03/using-golang-and-graphviz-to-visualize-complex-grails-applications/


* Go Distribution notes
    * Binary is in the terraform module

* Module packge notes
    * tfgraph go binary
    * aws icons
    * graphviz ?
* Misc links
    * display image in dot file layout: https://stackoverflow.com/questions/8073115/how-do-i-get-dot-to-display-an-image-for-a-node
    * online svg editor: http://www.drawsvg.org/
    * styling SVG with CSS: http://svgtutorial.com/styling-svg-with-css/
    * Go by Example: https://gobyexample.com/
        * Command-Line Flags: https://gobyexample.com/command-line-flags
        * Spawning Processes: https://gobyexample.com/spawning-processes
        * Reading Files: https://gobyexample.com/reading-files
        * Writing Files: https://gobyexample.com/writing-files
    * DOT File Links
        * https://www.graphviz.org/pdf/dotguide.pdf
    * Grails Dep Visualizer using Graphviz: https://github.com/ilikeorangutans/grails-service-visualizer
    * Go GraphViz Parser: https://github.com/awalterschulze/gographviz
        * GoDoc: https://godoc.org/github.com/awalterschulze/gographviz#SubGraph
    * Go Build Command: https://dave.cheney.net/2013/10/15/how-does-the-go-build-command-work

## Deps



## Building / Packaging
* pull gographviz dependency locally
    * `go get github.com/awalterschulze/gographviz`
* build tfgraph
    * `go build`
    * outputs tfgraph.exe
