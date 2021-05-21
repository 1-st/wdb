package cmd

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"unicode"
	"wdb/lib/serve"
	"wdb/lib/util"
)

const (
	EachLine = 4
	Size     = 144
)

func TrimLeftRight(str string) string {
	return strings.TrimRight(strings.TrimLeft(str, " "), " ")
}

func RunPrint(str string) {
	str = TrimLeftRight(str)
	args := strings.Split(str, " ")
	if len(args) != 2 {
		fmt.Println("print 需要2个参数")
		return
	}
	ptype := TrimLeftRight(args[0])
	pN, err := strconv.ParseInt(TrimLeftRight(args[1]), 10, 64)
	if err != nil {
		fmt.Println("print 第2个参数必须是数字")
		return
	}
	if ptype == "word" {
		PrintWord(int(pN))
	} else if ptype == "phrase" {
		PrintPhrase(int(pN))
	} else {
		fmt.Println("不支持的打印类型")
		return
	}
}

func PrintWord(n int) {
	var list util.PairList
	for _, v := range serve.DB.Cwords.Cword {
		if len(list) >= n {
			break
		}
		score := View2Score(v.Cviews)
		if v.Cok == nil || v.Cok.String == "false" {
			list = append(list, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
		}
	}
	list.RSort()
	var N = 0
	if n > len(list) {
		N = len(list)
	} else {
		N = n
	}

	var i = 0
	var file = ""
	for i < N {
		file += "\n\n"
		var eachLine = make([][]string, 0)
		for j := i; j < i+EachLine && j < N; j++ {
			ms := GetMeaningSimple(list[j].Name)
			if len(ms) > 0 {
				if !strings.HasPrefix(TrimLeftRight(ms[0]), "*") {
					ms = append([]string{""}, ms...)
				}
			}
			eachLine = append(eachLine, ms)
		}
		var maxX = 0
		for _, v := range eachLine {
			if len(v) > maxX {
				maxX = len(v)
			}
		}
		for x := 0; x < maxX; x++ {
			for y := 0; y < EachLine; y++ {
				if y < len(eachLine) {
					if x < len(eachLine[y]) {
						var num = ""
						if x == 1 {
							num = "①"
						} else if x == 2 {
							num = "②"
						} else {
							num = strconv.Itoa(i + 1 + y)
						}
						file += fillString(fmt.Sprintf("%s", num+eachLine[y][x]))
					} else {
						file += fillString(fmt.Sprintf("%s", "　　　　　　　　"))
					}
				} else {
					file += fillString(fmt.Sprintf("%s", "　　　　　　　　"))
				}
			}
			file += "\n"
		}
		i += EachLine
	}
	if err := ioutil.WriteFile(genFilename(), []byte(file), 0777); err != nil {
		fmt.Println("写入打印文件出错")
	}
}

func rLen(str string) int {
	var l = 0
	for _, v := range []rune(str) {
		if unicode.Is(unicode.Han, v) || unicode.IsUpper(v) {
			l += 2
		} else if unicode.IsSpace(v) {
			l += 1
		} else {
			l += 1
		}
	}
	return l
}

func fillString(str string) string {
	if rLen(str) < Size/EachLine {
		for rLen(str) < Size/EachLine {
			str += " "
		}
		return str
	} else {
		for rLen(str) >= Size/EachLine {
			words := []rune(str)
			words = words[:len(words)-1]
			str = string(words)
		}
		for rLen(str) < Size/EachLine {
			str += " "
		}
		return str
	}
}

func PrintPhrase(n int) {
	var list util.PairList
	var listExp util.PairList
	for _, v := range serve.DB.Cphrases.Cphrase {
		if len(list) >= n {
			break
		}
		score := View2Score(v.Cviews)
		if v.Cok == nil || v.Cok.String == "false" {
			list = append(list, util.Pair{
				Name:  v.Cid.String,
				Score: float32(score),
			})
			listExp = append(listExp, util.Pair{
				Name:  v.Cexplains.String,
				Score: float32(score),
			})
		}
	}
	list.RSort()
	listExp.RSort()
	var N = 0
	if n > len(list) {
		N = len(list)
	} else {
		N = n
	}

	var i = 0
	var file = ""
	for i < N {
		file += "\n"
		var eachLine = make([][]string, 0)
		for j := i; j < i+EachLine && j < N; j++ {
			ms := []string{listExp[i+j].Name}
			if len(ms) > 0 {
				if !strings.HasPrefix(TrimLeftRight(ms[0]), "*") {
					ms = append([]string{""}, ms...)
				}
			}
			eachLine = append(eachLine, ms)
		}
		var maxX = 0
		for _, v := range eachLine {
			if len(v) > maxX {
				maxX = len(v)
			}
		}
		for x := 0; x < maxX; x++ {
			for y := 0; y < EachLine; y++ {
				if y < len(eachLine) {
					if x < len(eachLine[y]) {
						var num = ""
						if x == 1 {
							num = ""
						} else if x == 2 {
							num = ""
						} else {
							num = strconv.Itoa(i + 1 + y)+"."
						}
						file += fillString(fmt.Sprintf("%s", num+eachLine[y][x]))
					} else {
						file += fillString(fmt.Sprintf("%s", "　　　　　　　　"))
					}
				} else {
					file += fillString(fmt.Sprintf("%s", "　　　　　　　　"))
				}
			}
			file += "\n"
		}
		i += EachLine
	}
	if err := ioutil.WriteFile(genFilename(), []byte(file), 0777); err != nil {
		fmt.Println("写入打印文件出错")
	}
}

func genFilename() string {
	MD5Str := strings.Split(strings.ReplaceAll(time.Now().String(), " ", "-"), ".")[0]
	return "print@" + MD5Str + ".txt"
}
