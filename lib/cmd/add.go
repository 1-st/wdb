package cmd

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
	"wdb/lib/constant"
	"wdb/lib/serve"
)

func RunAdd(str string) {
	for i := 0; i < len(str); i++ {
		if str[i] == '|' {
			fmt.Println("添加词组")
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
			return
		}
	}
	word := strings.Trim(str, " ")
	fmt.Println("添加单词")
	found := false
	for _, v := range serve.DB.Cwords.Cword {
		if v.Cid.String == word {
			found = true
		}
	}
	if found {
		fmt.Println("单词已经存在")
		return
	}
	AddWord(word)
}

func AddWord(word string) {
	t := serve.Dict.Translate(word)
	if t == nil {
		fmt.Println("词典中找不到单词!")
		return
	}
	fmt.Println("词典释义:")
	for _, v := range t {
		for _, vv := range v.Parts {
			fmt.Println(string(vv.Data))
		}
	}
	serve.DB.Cwords.Cword = append(serve.DB.Cwords.Cword, &constant.Cword{
		Cid: &constant.Cid{
			String: word,
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
	})
	if serve.SaveDB() {
		fmt.Println("更新数据库成功!")
	} else {
		fmt.Println("更新数据库失败!")
	}
}

func AddPhrase(line string, idx int) {
	var phrase = ""
	var explain = ""
	for i := 0; i < len(line); i++ {
		if line[i] == '|' {
			phrase = line[0:i]
			explain = line[i+1 : len(line)]
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
