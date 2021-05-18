package cmd

import (
	"fmt"
	"io"
	"wdb/lib/serve"
)

func RunList(str string,out io.Writer){
	fmt.Println(serve.DB)
}