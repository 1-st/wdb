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

func RunReviewR(str string) {
	var N int
	var cN int
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
	cN = N
	var list util.PairList
	for _, v := range serve.DB.Cwords.Cword {
		score := View2Score(v.Cviews)
		if v.Cok == nil || v.Cok.String == "false" {
			list = append(list, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		}
	}
	list.RSort()
	var i = len(list) - 1
	var cache []int
	var round = 1
	fmt.Println("\n反向\n")

	for N > 0 && i >= 0 || len(cache) != 0 {
		if len(cache) == 15 || !(N > 0 && i >= 0) {
			round = 2
		} else if len(cache) == 0 {
			round = 1
		}
		var ii int
		if round == 1 {
			ii = i
		} else {
			ii = cache[0]
		}

		fmt.Println(list[ii].Name)
		fmt.Printf("(%v / %v) round:%v\n", len(list)-1-ii+1, cN, round)
		fmt.Println("是否认识这个单词?(y/n, default is yes)")
		s := prompt.Input("review > ", completer)
		idx := 0
		for k, v := range serve.DB.Cwords.Cword {
			if list != nil && list[ii].Name == v.Cid.String {
				idx = k
			}
		}
		if s == "yes" || s == "y" || s == "" {
			views := serve.DB.Cwords.Cword[idx].Cviews
			if len(views.Cat) >= 3 && check12Hour(views.Cat[len(views.Cat)-1].String) {
				ok := serve.DB.Cwords.Cword[idx].Cok
				if ok == nil {
					ok = new(constant.Cok)
				}
				ok.String = "true"
			}
			AddView(views)
			AfterEnter(serve.DB.Cwords.Cword[idx].Cid.String)
		} else if s == "no" || s == "n" {
			serve.DB.Cwords.Cword[idx].Cviews.Cat = serve.DB.Cwords.Cword[idx].Cviews.Cat[:0]
			AfterEnter(serve.DB.Cwords.Cword[idx].Cid.String)
		} else if s == "exit" || s == "quit" || s == "q" {
			break
		}
		if round == 2 {
			cache = cache[1:]
		} else { //rand==1
			cache = append(cache, i)
			i--
			N--
		}
	}
	fmt.Printf("总共复习了 %v 个单词\n", len(list)-1-i)
}
