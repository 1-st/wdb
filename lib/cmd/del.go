package cmd

import (
	"fmt"
	"strings"
	"wdb/lib/serve"
)

func RunDel(str string){
	word := strings.Trim(str, " ")
	if strings.Contains(word, " ") {
		DelPhrase(word)
	} else {
		DelWord(word)
	}
}


func DelPhrase(phrase string){
	found := false
	idx:=0
	for k, v := range serve.DB.Cphrases.Cphrase {
		if v.Cid.String == phrase {
			found = true
			idx = k
		}
	}
	if !found {
		fmt.Println("词组不存在")
		return
	}else{
		serve.DB.Cphrases.Cphrase = append(serve.DB.Cphrases.Cphrase[:idx], serve.DB.Cphrases.Cphrase[idx+1:]...)
	}
	if serve.SaveDB(){
		fmt.Println("删除成功!")
	}else{
		fmt.Println("删除失败!")
	}
}

func DelWord(word string){
	found := false
	idx:=0
	for k, v := range serve.DB.Cwords.Cword {
		if v.Cid.String == word {
			found = true
			idx = k
		}
	}
	if !found {
		fmt.Println("单词不存在")
		return
	}else{
		serve.DB.Cwords.Cword = append(serve.DB.Cwords.Cword[:idx], serve.DB.Cwords.Cword[idx+1:]...)
	}
	if serve.SaveDB(){
		fmt.Println("删除成功!")
	}else{
		fmt.Println("删除失败!")
	}
}