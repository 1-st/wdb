package main

import (
	"log"
	"os"
	"wdb/lib"
)

func main() {
	if len(os.Args)==1 {
		lib.Execute()
	}else{
		log.Fatal("暂不支持启动参数")
	}
}
