package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"time"
	"wdb/lib/constant"
	"wdb/lib/serve"
	"wdb/lib/similar"
	"wdb/lib/similar/ai"
	"wdb/lib/util"
)

func FindWord(word string) {
	found := false
	fmt.Println()
	fmt.Println(word)
	fmt.Println()
	for _, v := range serve.DB.Cwords.Cword {
		if v.Cid.String == word {
			found = true
			color.Blue(GetMeaning(word))
			if v.Cviews == nil {
				v.Cviews = new(constant.Cviews)
			}
			PrintMemory(v.Cviews)
			AddView(v.Cviews)
			serve.SaveDB()

			//形似词
			var list = new(util.PairList)
			for _, v := range serve.DB.Cwords.Cword {
				if v.Cid.String != word {
					score, err := similar.GetDiffSimilarity(word, v.Cid.String)
					if err != nil {
						fmt.Println("diff出错")
					}
					to, err := strconv.ParseFloat(serve.ConfigBody.Cconfig.Csimilar_dash_word_dash_threshold_dash_diff.String, 32)
					if err != nil {
						fmt.Println("config diff 数值不是float")
						break
					}
					if score >= float32(to) {
						*list = append(*list, util.Pair{
							Name:  v.Cid.String,
							Score: score,
						})
					}
				}
			}
			list.Sort()
			if len(*list) != 0 {
				fmt.Println()
				fmt.Println("形似词:")
				for _, v := range *list {
					fmt.Printf("%v %v \t", v.Name, util.FtoS(float64(v.Score))+"%")
				}
				fmt.Println()
			}

			if serve.Model != nil {
				//近义词
				var list = new(util.PairList)
				for _, v := range serve.DB.Cwords.Cword {
					if v.Cid.String != word {
						score, err := ai.GetNetworkSimilarity(word, v.Cid.String)
						if err != nil {
							fmt.Println("word2vec model 出错")
						}
						to, err := strconv.ParseFloat(serve.ConfigBody.Cconfig.Csimilar_dash_word_dash_threshold_dash_network.String, 32)
						if err != nil {
							fmt.Println("config network 数值不是float")
							break
						}
						if score >= float32(to) {
							*list = append(*list, util.Pair{
								Name:  v.Cid.String,
								Score: score,
							})
						}
					}
				}
				list.Sort()
				if len(*list) != 0 {
					fmt.Println()
					fmt.Println("近义词:")
					for _, v := range *list {
						fmt.Printf("%v %v \t", v.Name, util.FtoS(float64(v.Score))+"%")
					}
					fmt.Println()
				}
				//网络近义词
				var N = 3
				fmt.Println("网络近义词:")
				match := ai.GetSimilarity(word, N)
				for i := 0; i < N; i++ {
					fmt.Printf("%v %v\t", match[i].Word, util.FtoS(float64(match[i].Score))+"%")
				}
				fmt.Println()

			}
			break
		}
	}
	if !found {
		fmt.Println("数据库中没有此单词")
		if m := GetMeaning(word); m != "" {
			fmt.Println()
			fmt.Println(word)
			fmt.Println()
			fmt.Println(m)
			fmt.Println()
		}
	}
}

func FindPhrase(phrase string) {
	found := false
	fmt.Println()
	fmt.Println(phrase)
	fmt.Println()
	for _, v := range serve.DB.Cphrases.Cphrase {
		if v.Cid.String == phrase {
			found = true
			color.Blue("%v\n", v.Cexplains.String)
			if v.Cviews == nil {
				v.Cviews = new(constant.Cviews)
			}
			PrintMemory(v.Cviews)
			AddView(v.Cviews)
			serve.SaveDB()
			break
		}
	}
	if !found {
		fmt.Println("数据库中没有此词组")
	}
}

func GetMeaning(word string) string {
	t := serve.Dict.Translate(word)
	if t == nil {
		return ""
	}
	for _, v := range t {
		for _, vv := range v.Parts {
			return string(vv.Data)
		}
	}
	return ""
}

func PrintMemory(view *constant.Cviews) {
	t := len(view.Cat)
	fmt.Println()
	fmt.Println("复习次数: " + strconv.Itoa(t))
	if t != 0 {
		lastTime := view.Cat[len(view.Cat)-1].String
		fmt.Println("上一次复习时间: " + lastTime)
		fmt.Printf("距今小时数: %v\n", strconv.FormatFloat(util.HourToNow(lastTime), 'f', 1, 64))
	}
	fmt.Println("预期记忆率: " + strconv.FormatFloat(View2Score(view), 'f', 1, 64) + "%")
}

func AddView(view *constant.Cviews) {
	view.Cat = append(view.Cat, &constant.Cat{
		String: time.Now().Format(time.RFC3339),
	})
}

func View2Score(view *constant.Cviews) float64 {
	var times []time.Time
	for _, v := range view.Cat {
		t, err := time.Parse(time.RFC3339, v.String)
		if err != nil {
			fmt.Println("读取时间出错: " + err.Error())
		}
		times = append(times, t)
	}
	return util.GetScore(times) * 100
}
