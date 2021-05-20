package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"wdb/lib/serve"
	"wdb/lib/util"
)

func RunList(str string) {
	//单词
	var list util.PairList
	var listOK util.PairList
	for _, v := range serve.DB.Cwords.Cword {
		score := View2Score(v.Cviews)
		if v.Cok==nil||v.Cok.String == "false" {
			list = append(list, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		} else {
			listOK = append(list, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		}
	}
	list.Sort()
	PrintList(&list)
	fmt.Println()
	fmt.Println()

	//词组
	var phrases util.PairList
	var phrasesOK util.PairList
	for _, v := range serve.DB.Cphrases.Cphrase {
		score := View2Score(v.Cviews)
		if v.Cok == nil || v.Cok.String == "false" {
			phrases = append(phrases, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		} else {
			phrasesOK = append(phrasesOK, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		}
	}
	PrintList(&phrases)
	fmt.Println()
	fmt.Println()

	//已完成
	if len(listOK)!=0||len(phrasesOK)!=0{
		fmt.Println("已完成:")
		for _, v := range listOK {
			color.Set(color.BgHiGreen).Print(" ")
			fmt.Printf("%v", v.Name)
			fmt.Print("\t")
		}
		for _, v := range phrasesOK {
			color.Set(color.BgHiGreen).Print(" ")
			fmt.Printf("%v", v.Name)
			fmt.Print("\t")
		}
		fmt.Println()
	}

	fmt.Printf("单词进度 %v/%v\n", len(listOK), len(list))
	fmt.Printf("词组进度 %v/%v\n", len(phrasesOK), len(phrases))
}

func PrintList(list *util.PairList){
	for _, v := range *list {
		if v.Score <= 30 {
			color.Set(color.BgHiRed).Print(" ")
		} else if v.Score <= 40 {
			color.Set(color.BgRed).Print(" ")
		} else if v.Score <= 50 {
			color.Set(color.BgHiYellow).Print(" ")
		} else if v.Score <= 60 {
			color.Set(color.BgYellow).Print(" ")
		} else if v.Score <= 85 {
			color.Set(color.BgGreen).Print(" ")
		} else {
			color.Set(color.BgHiGreen).Print(" ")
		}
		fmt.Printf("%v", v.Name)
		fmt.Print("\t")
	}
}
