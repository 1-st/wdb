package lib

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"wdb/lib/serve"
)

func Execute() {
	serve.PrintLogo()
	color.Blue("v0.1")
	for {
		t := prompt.Input("wdb > ", completer)
		if t == "" {
			continue
		}
		if t == "quit" || t == "exit" || t == "q" {
			break
		}
		i := 0
		for i < len(t) && t[i] != ' ' {
			i++
		}
		cmd := t[0:i]
		if v, ok := CmdMap[cmd]; ok {
			if i == len(t) || i == len(t)-1 {
				v("")
			} else {
				v(t[i+1:])
			}
		} else {
			fmt.Println("没有这个指令")
		}
	}
}
