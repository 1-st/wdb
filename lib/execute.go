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
		t := prompt.Input("mdb >> ", completer)
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
		if v, ok := CmdMap[t[0:i]]; ok {
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

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "list", Description: "打印记录在数据库中的单词和词组"},
		{Text: "help", Description: "打印帮助"},
		{Text: "find", Description: "[word/phrase] 寻找记录的单词或者词组"},
		{Text: "add", Description: "[word] [explain] 添加一个单词,如果已经存在则添加一项释义"},
		{Text: "addP", Description: "[phrase]<[explain]>添加一个词组，在词组后的尖括号内写上释义"},
		{Text: "del", Description: "[word/phrase] 删除单词或者词组"},
		{Text: "delX", Description: "[word][N] 删除单词或者词组第N个解释"},
		{Text: "review", Description: "[N] 复习单词,后接复习的数量(默认为3)"},
		{Text: "reviewP", Description: "[N] 复习词组,后接复习的数量(默认为3)"},

		{Text: "l", Description: "list"},
		{Text: "h", Description: "help"},
		{Text: "f", Description: "find"},
		{Text: "a", Description: "add"},
		{Text: "ap", Description: "addP"},
		{Text: "d", Description: "del"},
		{Text: "dx", Description: "delX"},
		{Text: "r", Description: "review"},
		{Text: "rp", Description: "reviewP"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
