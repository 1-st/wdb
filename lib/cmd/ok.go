package cmd

import (
	"fmt"
	"strings"
	"wdb/lib/constant"
	"wdb/lib/serve"
)

func RunOK(str string) {
	line := strings.TrimRight(strings.TrimLeft(str, " "), " ")
	if strings.Contains(line, " ") {
		PhraseOK(line)
	} else {
		WordOK(line)
	}
}

func WordOK(word string) {
	var found = false
	for _, v := range serve.DB.Cwords.Cword {
		if v.Cid.String == word {
			found = true
			if v.Cok == nil {
				v.Cok = new(constant.Cok)
				v.Cok.String = "true"
			} else {
				if v.Cok.String == "false" {
					v.Cok.String = "true"
				} else {
					v.Cok.String = "false"
				}
			}
			serve.SaveDB()
		}
	}
	if !found {
		fmt.Println("找不到单词: " + word)
	}
}

func PhraseOK(phrase string) {
	var found = false
	for _, v := range serve.DB.Cphrases.Cphrase {
		if v.Cid.String == phrase {
			found = true
			if v.Cok == nil {
				v.Cok = new(constant.Cok)
				v.Cok.String = "true"
			} else {
				if v.Cok.String == "false" {
					v.Cok.String = "true"
				} else {
					v.Cok.String = "false"
				}
			}
			serve.SaveDB()
		}
	}
	if !found {
		fmt.Println("找不到词组: " + phrase)
	}
}
