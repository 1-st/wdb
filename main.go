package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"wdb/lib"
	"wdb/lib/cmd"
	"wdb/lib/similar/ai"
)

func main() {
	if len(os.Args) != 1 {
		for _,v:= range os.Args[1:]{
			if v == "fast" {
				ai.FastBoot = true
			}else if strings.HasPrefix(v,"line="){
				var err error
				cmd.LineWord,err = strconv.ParseInt(strings.TrimPrefix(v,"line="),10,64)
				if err!=nil{
					fmt.Println("line=???")
				}
			}
		}
	}
	ai.Load()
	lib.Execute()
}
