package cmd

import (
	"io"
	"wdb/lib/serve"
)

func  RunHelp(str string,out io.Writer){
	serve.PrintHelp(out)
}
