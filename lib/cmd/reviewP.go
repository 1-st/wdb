package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"strconv"
	"strings"
	"wdb/lib/constant"
	"wdb/lib/serve"
	"wdb/lib/util"
)

func RunReviewP(str string) {
	var N int
	var err error
	if strings.Trim(str, " ") == "" {
		N = 30
	} else {
		N, err = strconv.Atoi(str)
		if err != nil {
			fmt.Println("参数错误")
			return
		}
	}
	var list util.PairList
	for _, v := range serve.DB.Cphrases.Cphrase {
		score := View2Score(v.Cviews)
		if v.Cok == nil || v.Cok.String == "false" {
			list = append(list, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		}
	}
	list.Sort()
	var i = len(list) - 1
	for N > 0 && i >= 0 {
		fmt.Println(list[i].Name)
		fmt.Println("是否认识这个词组?(y/n, default is yes)")
		s := prompt.Input("preview > ", completer)
		idx := 0
		for k, v := range serve.DB.Cphrases.Cphrase {
			if list != nil && list[i].Name == v.Cid.String {
				idx = k
			}
		}
		if s == "yes" || s == "y" || s == "" {
			views:=serve.DB.Cphrases.Cphrase[idx].Cviews
			if len(views.Cat)>=3&&check12Hour(views.Cat[len(views.Cat)-1].String){
				ok := serve.DB.Cphrases.Cphrase[idx].Cok
				if ok == nil {
					ok = new(constant.Cok)
				}
				ok.String = "true"
			}
			AddView(views)
			fmt.Println(serve.DB.Cphrases.Cphrase[idx].Cexplains.String)
			for i:=0;i<5;i++{
				fmt.Println()
			}
			serve.SaveDB()
		} else if s == "no" || s == "n" {
			serve.DB.Cphrases.Cphrase[idx].Cviews.Cat = serve.DB.Cphrases.Cphrase[idx].Cviews.Cat[:0]
			for i:=0;i<5;i++{
				fmt.Println()
			}
			serve.SaveDB()
		} else if s == "exit" || s == "quit" || s == "q" {
			break
		}
		i--
		N--
	}
	fmt.Printf("总共复习了 %v 个词组\n",len(list)-1-i)
}
