package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"wdb/lib/serve"
	"wdb/lib/util"
)

var LineWord int64 = 6

func RunList(str string) {
	line := strings.TrimLeft(strings.TrimRight(str, " "), " ")
	if line != "" {
		if strings.Contains(line, " ") {
			FindPhrase(line)
		} else {
			FindWord(line)
		}
	} else {
		List()
	}
}

func List() {
	fmt.Println()
	fmt.Println("单词:")
	fmt.Println()
	printClockInit()
	//单词
	var list util.PairList
	var listOK util.PairList
	for _, v := range serve.DB.Cwords.Cword {
		score := View2Score(v.Cviews)
		if v.Cok == nil || v.Cok.String == "false" {
			list = append(list, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		} else {
			listOK = append(listOK, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		}
	}
	list.Sort()
	PrintList(&list)
	fmt.Println()
	fmt.Println()
	fmt.Println("词组：")
	fmt.Println()

	printClockInit()
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

	//已完成
	if len(listOK) != 0 || len(phrasesOK) != 0 {
		if len(listOK) != 0 {
			fmt.Println()
			fmt.Println("已完成单词:")
			fmt.Println()
			printClockInit()
			for _, v := range listOK {
				printClock(color.Set(color.FgHiGreen).Sprint("●") + v.Name)
			}
			fmt.Println()
		}
		if len(phrasesOK) != 0 {
			fmt.Println()
			fmt.Println("已完成词组:")
			fmt.Println()
			printClock2Init()
			for _, v := range phrasesOK {
				printClock2(color.Set(color.FgHiGreen).Sprint("●") + v.Name)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Printf("单词进度: %v/%v\n", len(listOK), len(list)+len(listOK))
	fmt.Printf("词组进度: %v/%v\n", len(phrasesOK), len(phrases)+len(phrasesOK))
	fmt.Println()
	PrintPoint()
	fmt.Println()
}

func PrintPoint() {
	var f = func(colors ...color.Attribute) (res string) {
		for _, v := range colors {
			res += color.Set(v).Sprint("●")
		}
		return
	}
	fmt.Println(f(color.FgHiGreen, color.FgGreen, color.FgHiBlue,
		color.FgBlue, color.FgYellow, color.FgHiYellow,
		color.FgMagenta, color.FgHiMagenta, color.FgRed,
		color.FgHiRed, color.FgWhite))
}

func PrintList(list *util.PairList) {
	var point = ""
	for _, v := range *list {
		if v.Score <= 13 {
			point = color.Set(color.FgWhite).Sprint("●")
		} else if v.Score <= 15 {
			point = color.Set(color.FgHiRed).Sprint("●")
		} else if v.Score <= 17 {
			point = color.Set(color.FgRed).Sprint("●")
		} else if v.Score <= 19 {
			point = color.Set(color.FgHiMagenta).Sprint("●")
		} else if v.Score <= 22 {
			point = color.Set(color.FgMagenta).Sprint("●")
		} else if v.Score <= 25 {
			point = color.Set(color.FgHiYellow).Sprint("●")
		} else if v.Score <= 34 {
			point = color.Set(color.FgYellow).Sprint("●")
		} else if v.Score <= 38 {
			point = color.Set(color.FgBlue).Sprint("●")
		} else if v.Score <= 42 {
			point = color.Set(color.FgHiBlue).Sprint("●")
		} else if v.Score <= 60 {
			point = color.Set(color.FgGreen).Sprint("●")
		} else {
			point = color.Set(color.FgHiGreen).Sprint("●")
		}
		printClock(point + v.Name)
	}
}

var print_count = 1

func
printClockInit() {
	print_count = 1
}

func printClock(str string) {
	fmt.Printf("%-25s", str)
	if print_count%int(LineWord) == 0 {
		fmt.Print("\n")
	}
	print_count++
}

var print_count2 = 1

func printClock2Init() {
	print_count2 = 1
}
func printClock2(str string) {
	fmt.Printf("%-45s", str)
	if print_count%int(LineWord/2) == 0 {
		fmt.Print("\n")
	}
	print_count2++
}
