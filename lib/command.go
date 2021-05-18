package lib

import (
	"io"
	"os"
	"wdb/lib/cmd"
)

type RunFunc func(str string, out io.Writer)

var WrapStdout = func(runFunc RunFunc) func(str string) {
	return func(str string) {
		runFunc(str, os.Stdout)
	}
}

// TODO
//var WrapHttp = func(str string, f RunFunc) {
//
//}

var CmdMap = make(map[string]func(str string))

//var ApiMap = make(map[string]func(str string))

func init() {
	var add = func(m *map[string]func(str string), f func(str string), ss ...string) {
		for _, v := range ss {
			(*m)[v] = f
		}
	}
	add(&CmdMap, WrapStdout(cmd.RunAdd), "add", "a")
	add(&CmdMap, WrapStdout(cmd.RunAddP), "addP", "ap")
	add(&CmdMap, WrapStdout(cmd.RunDel), "del", "d")
	add(&CmdMap, WrapStdout(cmd.RunDelX), "delX", "dx")
	add(&CmdMap, WrapStdout(cmd.RunFind), "find", "f")
	add(&CmdMap, WrapStdout(cmd.RunHelp), "help", "h")
	add(&CmdMap, WrapStdout(cmd.RunList), "list", "l")
	add(&CmdMap, WrapStdout(cmd.RunReview), "review", "r")
	add(&CmdMap, WrapStdout(cmd.RunReviewP), "reviewP", "rp")
}
