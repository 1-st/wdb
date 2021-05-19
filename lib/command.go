package lib

import (
	"github.com/c-bata/go-prompt"
	"wdb/lib/cmd"
	"wdb/lib/constant"
	"wdb/lib/serve"
)

var CmdMap = make(map[string]func(str string))

func init() {
	var add = func(m *map[string]func(string), f func(string), ss ...string) {
		for _, v := range ss {
			(*m)[v] = f
		}
	}
	add(&CmdMap, cmd.RunAdd, "add", "a")
	add(&CmdMap, cmd.RunAddP, "addP", "ap")
	add(&CmdMap, cmd.RunDel, "del", "d")
	add(&CmdMap, cmd.RunFind, "find", "f")
	add(&CmdMap, cmd.RunHelp, "help", "h")
	add(&CmdMap, cmd.RunList, "list", "l", "ls")
	add(&CmdMap, cmd.RunReview, "review", "r")
	add(&CmdMap, cmd.RunReviewP, "reviewP", "rp")
	add(&CmdMap, cmd.RunOK, "ok")
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "list", Description: constant.CMD_LIST_DSP},
		{Text: "help", Description: constant.CMD_HELP_DSP},
		{Text: "find", Description: constant.CMD_FIND_DSP},
		{Text: "add", Description: constant.CMD_ADD_DSP},
		{Text: "addP", Description: constant.CMD_ADDP_DSP},
		{Text: "del", Description: constant.CMD_DEL_DSP},
		{Text: "review", Description: constant.CMD_REVIEW_DSP},
		{Text: "reviewP", Description: constant.CMD_REVIEWP_DSP},
		{Text: "ok", Description: constant.CMD_OK_DSP},

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
	for _, v := range serve.DB.Cwords.Cword {
		s = append(s, prompt.Suggest{
			Text:        "find " + v.Cid.String,
			Description: "",
		})
	}
	for _, v := range serve.DB.Cphrases.Cphrase {
		s = append(s, prompt.Suggest{
			Text:        "find " + v.Cid.String,
			Description: "",
		})
	}
	return prompt.FilterContains(s, d.GetWordBeforeCursorWithSpace(), true)
}