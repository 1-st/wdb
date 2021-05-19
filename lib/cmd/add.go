package cmd

import (
	"fmt"
	"strings"
	"time"
	"wdb/lib/constant"
	"wdb/lib/serve"
)

func RunAdd(str string) {
	word := strings.Trim(str, " ")

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
