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
	add(&CmdMap, cmd.RunDel, "rm", "del")
	add(&CmdMap, cmd.RunHelp, "help", "h")
	add(&CmdMap, cmd.RunList, "list", "l", "ls")
	add(&CmdMap, cmd.RunReview, "review", "r")
	add(&CmdMap, cmd.RunReviewP, "reviewP", "rp")
	add(&CmdMap, cmd.RunReviewR, "reviewR", "rr")

	add(&CmdMap, cmd.RunOK, "ok")
	add(&CmdMap, cmd.RunPrint, "print", "p")
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "ls", Description: constant.CMD_LIST_DSP},
		{Text: "help", Description: constant.CMD_HELP_DSP},
		{Text: "add", Description: constant.CMD_ADD_DSP},
		{Text: "rm", Description: constant.CMD_RM_DSP},
		{Text: "review", Description: constant.CMD_REVIEW_DSP},
		{Text: "reviewP", Description: constant.CMD_REVIEWP_DSP},
		{Text: "ok", Description: constant.CMD_OK_DSP},
		{Text: "print", Description: constant.CMD_PRINT_DSP},

		{Text: "h", Description: "help"},
		{Text: "p", Description: "print"},
		{Text: "a", Description: "add"},
		{Text: "r", Description: "review"},
		{Text: "rp", Description: "reviewP"},
		{Text: "rr", Description: "reviewR"},
	}
	for _, v := range serve.DB.Cwords.Cword {
		s = append(s, prompt.Suggest{
			Text:        "ls " + v.Cid.String,
			Description: "",
		})
	}
	for _, v := range serve.DB.Cphrases.Cphrase {
		s = append(s, prompt.Suggest{
			Text:        "ls " + v.Cid.String,
			Description: "",
		})
	}
	return prompt.FilterFuzzy(s, d.GetWordBeforeCursorWithSpace(), true)
}
