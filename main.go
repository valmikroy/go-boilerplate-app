package main

import "go-boilerplate-app/cmd"

var VERSION string = "0.0.1-default"

func main() {
	cmd.Version = VERSION
	cmd.Execute()
}
