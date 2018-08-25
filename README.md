# tfgraph
* Generate AWS-Style graphs of Terraform infrastructure on AWS
* Written in Go, distributed as a terraform module
* Simple commandline to export a graph
* Runs on Windows, Mac and Linux

## Add to your Terraform project
* Add the module reference to your terraform `main.tf` (latest version on github):  
`module "tfgraph" {`  
`source = github.com/donbecker/tfgraph.git`  
`}`  

* Add the module reference to your terraform `main.tf` (specific version on github):  
`module "tfgraph" {`  
`source = github.com/donbecker/tfgraph.git?ref=v1.0.0`  
`}`  

* Add the module reference to your terraform `main.tf` (from your s3 bucket):  
`module "tfgraph" {`  
`source = "yours3bucket.s3-us-east-1.amazonaws.com/donbecker.tfgraph-v1.0.0.tar.gz"`  
`}` 

## Create your graph
* `terraform init` (to retrieve the tfgraph module)
* `tfgraph .`

