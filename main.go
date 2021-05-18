package main

import (
	"os"
	"wdb/lib"
)

func main() {
	if len(os.Args)==1 {
		lib.Execute()
	}
	//TODO http
}
