package cmd

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
	"wdb/lib/constant"
	"wdb/lib/serve"
)

func RunAddP(str string) {
	line := strings.TrimRight(str, " ")

	found := false
	idx := -1
	for k, v := range serve.DB.Cphrases.Cphrase {
		if v.Cid.String == line {
			found = true
			idx = k
		}
	}
	if found {
		fmt.Println("词组已经存在，修改释义")
		AddPhrase(strings.TrimLeft(line, " "), idx)
	}
	AddPhrase(strings.TrimLeft(line, " "), -1)
}

func AddPhrase(line string, idx int) {
	var phrase = ""
	var explain = ""
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == '(' {
			phrase = line[0:i]
			explain = line[i+1 : len(line)-1]
		}
	}
	if idx == -1 {
		serve.DB.Cphrases.Cphrase = append(serve.DB.Cphrases.Cphrase, NewPhrase(phrase, explain))
	} else {
		serve.DB.Cphrases.Cphrase[idx].Cok.String = "true"
	}

	if serve.SaveDB() {
		fmt.Println("更新数据库成功!")
	} else {
		fmt.Println("更新数据库失败!")
	}
}

func NewPhrase(phrase string, explain string) *constant.Cphrase {
	return &constant.Cphrase{
		Cid: &constant.Cid{
			String: strings.TrimRight(phrase, " "),
		},
		Cok: &constant.Cok{
			String: "false",
		},
		Cviews: &constant.Cviews{
			Cat: []*constant.Cat{
				&constant.Cat{
					String: time.Now().Format(time.RFC3339),
				},
			},
		},
		Cexplains: &constant.Cexplains{
			XMLName: xml.Name{},
			String:  strings.TrimLeft(strings.TrimRight(explain, " "), " "),
		},
	}
}
